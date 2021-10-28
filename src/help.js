const Help = require("@oclif/plugin-help").default;
const chalk = require("chalk");
const { say } = require("cfonts");
const { sortBy, uniqBy } = require("../tools/bool");
const { renderList } = require("../tools/list");
const indent = require("indent-string");
const { CommandHelp } = require("../tools/help_command");
const root = require("../contents/helpers/root");
const { TopicFormatter } = require("../contents/helpers/topic");

// constants
const { PRIMARY_COLOR } = require("../constants");

function getHelpSubject(args) {
  for (const arg of args) {
    if (arg === "--") return;
    if (arg === "help" || arg === "--help" || arg === "-h") continue;
    if (arg.startsWith("-")) return;

    return arg;
  }
}

module.exports = class MyHelpClass extends Help {
  get sortedCommands() {
    let commands = this.config.commands;

    commands = commands.filter((c) => this.opts.all || !c.hidden);
    commands = sortBy(commands, (c) => c.id);
    commands = uniqBy(commands, (c) => c.id);

    console.log(commands);
  }

  get sortedTopics() {
    let topics = this._topics;
    topics = topics.filter((t) => this.opts.all || !t.hidden);
    topics = sortBy(topics, (t) => t.name);
    topics = uniqBy(topics, (t) => t.name);

    return topics;
  }

  showHelp(args) {
    // print secman cli version
    console.log(`${chalk.bold("Secman CLI")} ${this.config.version}\n`);

    const subject = getHelpSubject(args);

    if (!subject) {
      const rootCmd = this.config.findCommand("");
      if (rootCmd) this.showCommandHelp(rootCmd);
      this.showRootHelp();
      return;
    }

    const command = this.config.findCommand(subject);
    if (command) {
      this.showCommandHelp(command);
      return;
    }
  }

  showCommandHelp(command) {
    const name = command.id;
    const depth = name.split(":").length;

    const subTopics = this.sortedTopics.filter(
      (t) =>
        t.name.startsWith(name + ":") && t.name.split(":").length === depth + 1
    );

    const title =
      command.description && this.render(command.description).split("\n")[0];
    if (title) console.log(title + "\n");
    console.log(this.formatCommand(command));
    console.log("");

    if (subTopics.length > 0) {
      console.log(this.formatTopics(subTopics));
      console.log("");
    }

    root.learnMore(name);
  }

  async showRootHelp() {
    say("secman", {
      font: "3d",
      align: "left",
      colors: [PRIMARY_COLOR],
      background: "transparent",
    });

    console.log(chalk.grey.bold("USAGE\n"));
    console.log(chalk.cyan("  $ secman <COMMAND> [FLAGS]\n"));
    this.formatCommands(this.config.commands);
    root.root();
  }

  formatCommand(command) {
    const help = new CommandHelp(command, this.config, this.opts);
    return help.generate();
  }

  formatCommands(commands) {
    if (commands.length === 0) return "";

    const body = renderList(
      commands.map((c) => [
        c.id,
        c.description && this.render(c.description.split("\n")[0]),
      ]),
      {
        spacer: "\n",
        stripAnsi: this.opts.stripAnsi,
        maxWidth: this.opts.maxWidth - 2,
      }
    );

    console.log([chalk.grey.bold("COMMANDS\n"), indent(body, 2)].join("\n"));
  }

  formatTopic(topic) {
    TopicFormatter(this.render, topic, this.config, this.opts);
  }

  formatTopics(topics) {
    if (topics.length === 0) return "";
    const body = renderList(
      topics.map((c) => [
        c.name,
        c.description && this.render(c.description.split("\n")[0]),
      ]),
      {
        spacer: "\n",
        stripAnsi: this.opts.stripAnsi,
        maxWidth: this.opts.maxWidth - 2,
      }
    );

    return [chalk.grey.bold("TOPICS"), indent(body, 2)].join("\n");
  }
};
