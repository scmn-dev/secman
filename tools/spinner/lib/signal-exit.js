var process = global.process;

if (typeof process !== "object" || !process) {
  module.exports = function () {};
} else {
  var assert = require("assert");
  var signals = require("./signals");
  var isWin = /^win/i.test(process.platform);
  var EE = require("events");
  var emitter;
  var sigListeners = {};
  var loaded = false;

  if (typeof EE !== "function") {
    EE = EE.EventEmitter;
  }

  if (process.__signal_exit_emitter__) {
    emitter = process.__signal_exit_emitter__;
  } else {
    emitter = process.__signal_exit_emitter__ = new EE();
    emitter.count = 0;
    emitter.emitted = {};
  }

  if (!emitter.infinite) {
    emitter.setMaxListeners(Infinity);
    emitter.infinite = true;
  }

  module.exports = function (cb, opts) {
    if (global.process !== process) {
      return;
    }

    assert.equal(
      typeof cb,
      "function",
      "a callback must be provided for exit handler"
    );

    if (loaded === false) {
      load();
    }

    var ev = "exit";
    if (opts && opts.alwaysLast) {
      ev = "afterexit";
    }

    var remove = function () {
      emitter.removeListener(ev, cb);

      if (
        emitter.listeners("exit").length === 0 &&
        emitter.listeners("afterexit").length === 0
      ) {
        unload();
      }
    };

    emitter.on(ev, cb);

    return remove;
  };

  var unload = function unload() {
    if (!loaded || global.process !== process) {
      return;
    }

    loaded = false;

    signals.forEach(function (sig) {
      try {
        process.removeListener(sig, sigListeners[sig]);
      } catch (er) {}
    });

    process.emit = originalProcessEmit;
    process.reallyExit = originalProcessReallyExit;
    emitter.count -= 1;
  };

  module.exports.unload = unload;

  var emit = function emit(event, code, signal) {
    if (emitter.emitted[event]) {
      return;
    }

    emitter.emitted[event] = true;
    emitter.emit(event, code, signal);
  };

  signals.forEach(function (sig) {
    sigListeners[sig] = function listener() {
      if (process !== global.process) {
        return;
      }

      var listeners = process.listeners(sig);

      if (listeners.length === emitter.count) {
        unload();
        emit("exit", null, sig);
        emit("afterexit", null, sig);

        if (isWin && sig === "SIGHUP") {
          sig = "SIGINT";
        }

        process.kill(process.pid, sig);
      }
    };
  });

  module.exports.signals = function () {
    return signals;
  };

  var load = function load() {
    if (loaded || process !== global.process) {
      return;
    }

    loaded = true;

    emitter.count += 1;

    signals = signals.filter(function (sig) {
      try {
        process.on(sig, sigListeners[sig]);
        return true;
      } catch (er) {
        return false;
      }
    });

    process.emit = processEmit;
    process.reallyExit = processReallyExit;
  };

  module.exports.load = load;

  var originalProcessReallyExit = process.reallyExit;
  var processReallyExit = function processReallyExit(code) {
    if (process !== global.process) {
      return;
    }

    process.exitCode = code || 0;

    emit("exit", process.exitCode, null);
    emit("afterexit", process.exitCode, null);

    originalProcessReallyExit.call(process, process.exitCode);
  };

  var originalProcessEmit = process.emit;

  var processEmit = function processEmit(ev, arg) {
    if (ev === "exit" && process === global.process) {
      if (arg !== undefined) {
        process.exitCode = arg;
      }

      var ret = originalProcessEmit.apply(this, arguments);

      emit("exit", process.exitCode, null);
      emit("afterexit", process.exitCode, null);

      return ret;
    } else {
      return originalProcessEmit.apply(this, arguments);
    }
  };
}
