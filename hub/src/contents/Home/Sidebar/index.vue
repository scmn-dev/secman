<template>
  <div class="sidebar">
    <div class="account">
      <!-- Avatar -->
      <VAvatar :name="user.name" />
      <!-- Info -->
      <div class="account-info">
        <span class="account-info-name" v-text="user.name" />
      </div>

      <!-- Menu -->
      <button @click="showAccountMenu = !showAccountMenu">
        <VIcon
          :name="showAccountMenu ? 'x' : 'menu'"
          size="15px"
          class="c-white"
        />
      </button>

      <!-- Menu -->
      <div class="account-menu" :class="accountMenuClass">
        <div class="d-flex flex-column flex-items-start p-3">
          <!-- Logout -->
          <button class="c-danger mr-3" @click="onClickLogout">
            {{ $t("Logout") }}
            <VIcon name="logout" size="14px" rotation="180" class="ml-2" />
          </button>

          <div class="version">
            <p class="pt-3">
              Version:<strong>&nbsp;{{ _version() }}</strong>
            </p>
          </div>
        </div>
      </div>
    </div>

    <!-- Logins -->
    <MenuItem
      :service="$C.Services.Logins"
      :name="$t('Logins')"
      icon="lock-closed"
    />

    <!-- Credit Cards -->
    <MenuItem
      :service="$C.Services.CreditCards"
      :name="$t('Credit Cards')"
      icon="credit-card"
    />

    <!-- Emails -->
    <MenuItem :service="$C.Services.Emails" :name="$t('Emails')" icon="email" />

    <!-- Notes -->
    <MenuItem :service="$C.Services.Notes" :name="$t('Notes')" icon="note" />

    <!-- Servers -->
    <MenuItem
      :service="$C.Services.Servers"
      :name="$t('Servers')"
      icon="server"
    />

    <button class="btn-empty-fix"></button>

    <!-- products -->
    <button class="btn-products">
      <div class="icon">
        <a href="https://secman.dev" target="_blank" rel="noopener noreferrer">
          <VIcon
            v-tooltip="'View all Secman products'"
            name="infinity"
            size="14px"
          />
        </a>
      </div>
    </button>
  </div>
</template>

<script>
import MenuItem from "./MenuItem";
import { mapActions, mapState } from "vuex";

export default {
  components: {
    MenuItem,
  },

  data() {
    return {
      showAccountMenu: false,
    };
  },

  computed: {
    ...mapState(["user"]),

    accountMenuClass() {
      return { "--open": this.showAccountMenu };
    },
  },

  methods: {
    ...mapActions(["Logout"]),
    ...mapActions("Logins", ["FetchAll"]),

    onClickLogout() {
      this.Logout();
      this.$router.replace({ name: "Login" });
    },

    _version() {
      const { version } = require("../../../../package.json");
      return "v" + version;
    },
  },
};
</script>

<style lang="scss">
.version {
  border-top: 1px solid $color-gray-400;
  width: 140px;
}

.sidebar {
  width: 200px;
  min-width: 200px;
  height: 100%;
  display: flex;
  flex-direction: column;
  user-select: none;
  border-right: 1px solid $color-gray-400;

  .account {
    position: relative;
    margin: 32px 0px;
    width: 100%;
    height: 40px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0 $spacer-3;

    &-menu {
      position: absolute;
      background-color: $color-gray-500;
      border-radius: 12px;
      height: 0px;
      color: white;
      top: 55px;
      left: 12px;
      right: 12px;
      z-index: 2;
      transition: all 0.1s ease;
      overflow: hidden;

      &.--open {
        border: 1px solid $color-gray-400;
        height: 110px;
      }

      hr {
        width: 174px;
        border-bottom: 2px solid $color-gray-400;
        margin-bottom: 10px;
        margin-left: -16px;
      }

      button {
        color: white;
        font-size: $font-size-medium;
        margin-bottom: 20px;
        display: flex;
        align-items: center;

        &:last-child {
          margin-bottom: 0px;
        }
      }
    }

    &-info {
      display: flex;
      flex-direction: column;
      justify-content: space-between;
      margin: 0 auto 0 $spacer-2;

      &-name {
        font-size: 14px;
        line-height: 22px;
        color: #fff;
      }
    }
  }

  .btn-empty-fix {
    margin-top: auto;
  }

  .update-box {
    height: 30px;
    color: #fff;
    background-color: $color-primary;
    margin-top: auto;

    &:hover {
      opacity: 0.9;
    }
  }

  .btn-products {
    position: relative;
    height: 40px;
    background-color: $color-gray-500;
    font-size: $font-size-normal;
    border-top: 1px solid $color-gray-400;
    color: #fff;

    &,
    .icon {
      display: flex;
      align-items: center;
      justify-content: center;
    }

    .icon {
      width: 20px;
      height: 20px;
      border-radius: 4px;
      margin-left: $spacer-2;
      background-color: $color-gray-400;
    }

    .left-corner,
    .right-corner {
      position: absolute;
      z-index: 10;
    }

    .right-corner {
      color: $color-primary;
      top: 0px;
      right: 0px;
    }

    .left-corner {
      left: 0px;
      bottom: 0px;
      color: $color-secondary;
    }
  }

  .top-btn {
    color: $color-gray-300;
  }
  .top-btn:hover,
  .top-btn.active {
    color: $color-secondary;
  }

  &-menu-item {
    display: flex;
    align-items: center;
    padding-left: $spacer-3;
    height: 40px;
    color: $color-gray-300;
    font-size: $font-size-normal;
    border-top: 1px solid $color-gray-400;

    &.--lock {
      opacity: 0.6;
      cursor: not-allowed;
      pointer-events: none;
    }

    &:nth-last-child(3) {
      margin-bottom: auto;
    }

    &:last-of-type {
      border-bottom: 0;
      margin-bottom: 20px;
    }

    svg {
      color: $color-gray-300;
      margin-right: $spacer-2;
    }

    &.router-link-active {
      color: #fff;

      svg {
        color: $color-secondary;
      }
    }
  }
}
</style>
