<template>
  <div id="app">
    <header class="app-header">
      <button
        v-tooltip="`Secman Hub 📦`"
        class="trsn logo app-header-left-section"
      >
        <VIcon name="sm-logo" />
      </button>
    </header>
    <!-- Content -->
    <RouterView />
    <!-- Hidden -->
    <Icons />
    <notifications width="345" position="top center" />
  </div>
</template>

<script>
export default {
  data() {
    return {
      installEvent: undefined,
      shown: false,
    };
  },

  beforeMount() {
    window.addEventListener("beforeinstallprompt", (e) => {
      e.preventDefault();
      this.shown = true;
    });

    if (
      window.matchMedia("(display-mode: standalone)").matches ||
      window.navigator.standalone
    ) {
      this.shown = false;
    } else {
      this.shown = true;
    }
  },

  installPWA() {
    this.installEvent.prompt();
    // then refresh the page
    this.installEvent.userChoice.then((choiceResult) => {
      if (choiceResult.outcome === "accepted") {
        window.location.reload();
      }
    });
  },

  dismissPrompt() {
    this.shown = false;
  },
};
</script>

<style lang="scss">
.secman {
  font-weight: 600;
  font-size: $font-size-normal;
  color: $color-gray-200;
  line-height: 16px;
  padding-top: 4px;
}

.app-header {
  width: 100vw;
  height: 56px;
  margin: 0px;
  padding: 0px 50px;
  background: $color-gray-500;
  -webkit-app-region: drag;
  border-bottom: 1px solid $color-gray-400;

  & > * {
    -webkit-app-region: no-drag;
  }

  &,
  &-control-buttons {
    display: flex;
    align-items: center;
  }

  &-search {
    width: 295px;
    text-align: center;

    &-wrapper {
      position: relative;
      width: 260px;

      input {
        width: 100%;
        height: 40px;
        background-color: #000;
        border-radius: 8px;
        padding: 0 46px 0 24px;
        color: white;
        border: 0;
        font-size: 12px;

        &::placeholder {
          font-size: 12px;
          color: $color-gray-300;
        }
      }
    }

    .v-icon {
      top: 12px;
      right: 20px;
      position: absolute;
      color: $color-gray-300;
    }
  }

  &-control-buttons {
    width: 56px;
    height: 48px;
    -webkit-app-region: drag;
    margin-left: auto;
    display: flex;

    button {
      margin-right: 7px;
      -webkit-app-region: no-drag;
      color: $color-gray-300;
    }
  }

  &-left-section {
    display: flex;
    margin-right: auto;
  }
}
</style>
