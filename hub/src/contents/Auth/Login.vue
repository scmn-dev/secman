<template>
  <div class="login-container">
    <!-- Login Form -->
    <form class="login-form" @submit.stop.prevent="onLogin">
      <!-- E-Mail Address -->
      <div class="mt-4">
        <label v-text="$t('EMailAddress')" />
        <VFormText
          v-model="LoginForm.email"
          size="medium"
          v-validate="'required|email'"
          name="E-Mail"
          placeholder="E-Mail"
        />
      </div>
      <!-- Master Password -->
      <div class="mt-4 mb-5">
        <label class="w-100">
          {{ $t("MasterPassword") }}
        </label>
        <VFormText
          v-model="LoginForm.master_password"
          size="medium"
          type="password"
          name="Master Password"
          placeholder="Master Password"
          v-validate="'required|min:6|max:100'"
        />
      </div>

      <!-- Login Btn -->
      <VButton
        type="submit"
        :loading="$wait.is($waiters.Auth.Login)"
        size="medium"
        class="mb-2"
      >
        {{ $t("Login") }}
      </VButton>

      <VButton size="medium" style="background-color: white; color: black">
        <a @click="openPortal">
          {{ $t("CreateAccount") }}
        </a>
      </VButton>
    </form>
  </div>
</template>

<script>
import { mapActions } from "vuex";

export default {
  data() {
    return {
      installEvent: undefined,
      shown: false,
      windowRef: null,
      LoginForm: {
        email: "",
        master_password: "",
      },
    };
  },

  watch: {
    open(newOpen) {
      if (newOpen) {
        this.openPortal();
      } else {
        this.closePortal();
      }
    },
  },

  mounted() {
    if (this.open) {
      this.openPortal();
    }
  },

  beforeDestroy() {
    if (this.windowRef) {
      this.closePortal();
    }
  },

  beforeMount() {
    window.addEventListener("beforeinstallprompt", (e) => {
      e.preventDefault();
      this.installEvent = e;
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

  methods: {
    ...mapActions(["Login"]),

    openPortal() {
      this.windowRef = window.open(
        "https://auth.secman.dev",
        "",
        "width=500,height=600,left=200,top=200"
      );
      this.windowRef.addEventListener("beforeunload", this.closePortal);
    },

    closePortal() {
      if (this.windowRef) {
        this.windowRef.close();
        this.windowRef = null;
        this.$emit("close");
      }
    },

    onLogin() {
      this.$validator.validate().then(async (result) => {
        if (!result) return;

        const onError = (error) => {
          let text = this.$t("Ooops! Something went wrong! try again");
          if (error.response.status == 401) {
            text = this.$t(error.response.data.message);
          }

          this.$notifyError(text);
        };

        const onSuccess = async () => {
          await this.Login({ ...this.LoginForm });
          this.$router.replace({ name: "Home" });
        };

        this.$request(onSuccess, this.$waiters.Auth.Login, onError);
      });
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
  },
};
</script>

<style lang="scss">
.pwa {
  margin: 0;
  position: absolute;
  top: 50%;
  -ms-transform: translateY(-50%);
  transform: translateY(-50%);
}

.pwa-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 24px;
  height: 24px;
  border-radius: 4px;
  background-color: black;
  color: $color-gray-300;
}

.pwa-btn:hover {
  color: $color-secondary;
}

form {
  display: inline-block;
  margin-left: auto;
  margin-right: auto;
  text-align: left;
}

.login-container {
  width: 100%;
  height: calc(100% - 56px);
  position: absolute;
  display: flex;
  justify-content: space-between;
  background: url("./bg.gif");

  .btn,
  .form-text-wrapper {
    width: 350px;
  }
}

.master-pass-tooltip {
  color: $color-gray-300;
  float: right;
  margin-top: 4px;

  &:hover {
    color: #fff;
  }
}

.login-background {
  position: relative;
  top: 0px;
  left: 0px;
}

.login-right-corner {
  color: $color-primary;
  position: absolute;
  top: 0px;
  right: 0px;
  z-index: 10;
}

.login-left-corner {
  color: $color-primary;
  position: absolute;
  bottom: 0px;
  left: 0px;
  z-index: 10;
}

.login-form {
  width: 50%;
  min-width: 450px;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding-top: 60px;
  // background-color: black;
}

.login-form {
  label {
    display: block;
    font-size: 14px;
    font-weight: 500;
    line-height: 22px;
    color: #fff;
    margin-bottom: 8px;
  }
}
</style>
