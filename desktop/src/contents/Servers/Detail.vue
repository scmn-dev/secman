<template>
  <div class="detail-page">
    <div class="detail-page-header">
      <!-- Avatar -->
      <div class="detail-page-header-avatar">
        <CompanyLogo :url="form.url" />
      </div>
      <!-- Summary -->
      <div class="detail-page-header-summary">
        <span v-text="getTitle" class="url" />
        <span v-text="form.ip" class="email" />
      </div>

      <EditBtn v-if="!isEditMode" @click="isEditMode = $event" />
      <ClipboardButton :copy="copyContent" />
      <DeleteBtn @click="onClickDelete" />
    </div>
    <!-- Content -->
    <PerfectScrollbar class="detail-page-content">
      <form class="form" @submit.stop.prevent="onClickUpdate">
        <!-- Title -->
        <FormRowText
          v-model="form.title"
          :title="$t('TITLE')"
          :edit-mode="isEditMode"
        />

        <!-- IP -->
        <FormRowText
          v-model="form.ip"
          :title="$t('IP ADDRESS')"
          :edit-mode="isEditMode"
        />

        <!-- Username -->
        <FormRowText
          v-model="form.username"
          :title="$t('USERNAME')"
          :edit-mode="isEditMode"
        />

        <!-- Password -->
        <div class="form-row">
          <label v-text="$t('PASSWORD')" />
          <div class="d-flex">
            <VFormText
              v-if="isEditMode"
              :type="showPass ? 'text' : 'password'"
              v-model="form.password"
              :placeholder="$t('ClickToFill')"
              theme="no-border"
            />
            <!-- Text -->
            <div v-else class="d-flex flex-items-center px-3 py-2">
              <span
                v-text="showPass ? form.password : _password"
                class="mr-2"
              />
            </div>
            <GeneratePassword v-if="isEditMode" v-model="form.password" />
            <CheckPass :password="form.password" />
            <ShowPassBtn @click="showPass = $event" />
            <ClipboardButton
              v-if="form.password.length != 0"
              :copy="form.password"
            />
          </div>
        </div>
        <!-- URL -->
        <div class="form-row">
          <label v-text="$t('URL')" />
          <VFormText
            v-if="isEditMode"
            v-model="form.url"
            theme="no-border"
            :placeholder="$t('ClickToFill')"
          />
          <!-- Text -->
          <div v-else class="d-flex flex-items-center px-3 py-2">
            <span v-text="form.url" class="mr-2" />
            <ClipboardButton :copy="form.url" />
          </div>
        </div>
        <!-- HostingUsername -->
        <div class="form-row">
          <label v-text="$t('HOSTING USERNAME')" />
          <VFormText
            v-if="isEditMode"
            v-model="form.hosting_username"
            theme="no-border"
            :placeholder="$t('ClickToFill')"
          />
          <!-- Text -->
          <div v-else class="d-flex flex-items-center px-3 py-2">
            <span v-text="_hu" class="mr-2" />
            <ClipboardButton
              v-if="form.hosting_username != 0"
              :copy="form.hosting_username"
            />
          </div>
        </div>
        <!-- HostingPassword -->
        <div class="form-row">
          <label v-text="$t('HOSTING PASSWORD')" />
          <div class="d-flex">
            <VFormText
              v-if="isEditMode"
              :type="showHostingPass ? 'text' : 'password'"
              v-model="form.hosting_password"
              :placeholder="$t('ClickToFill')"
              theme="no-border"
            />
            <!-- Text -->
            <div v-else class="d-flex flex-items-center px-3 py-2">
              <span
                v-text="showHostingPass ? form.hosting_password : _hp"
                class="mr-2"
              />
            </div>
            <!-- Copy -->
            <GeneratePassword
              v-if="isEditMode"
              v-model="form.hosting_password"
            />
            <CheckPass :password="form.hosting_password" />
            <ShowPassBtn @click="showHostingPass = $event" />
            <ClipboardButton
              v-if="form.hosting_password.length != 0"
              :copy="form.hosting_password"
            />
          </div>
        </div>
        <!-- AdminUsername -->
        <div class="form-row">
          <label v-text="$t('ADMIN USERNAME')" />
          <VFormText
            v-if="isEditMode"
            v-model="form.admin_username"
            theme="no-border"
            :placeholder="$t('ClickToFill')"
          />
          <!-- Text -->
          <div v-else class="d-flex flex-items-center px-3 py-2">
            <span v-text="_au" class="mr-2" />
            <ClipboardButton
              v-if="form.admin_username.length != 0"
              :copy="form.admin_username"
            />
          </div>
        </div>
        <!-- AdminPassword -->
        <div class="form-row">
          <label v-text="$t('ADMIN PASSWORD')" />
          <div class="d-flex">
            <VFormText
              v-if="isEditMode"
              :type="showAdminPass ? 'text' : 'password'"
              v-model="form.admin_password"
              :placeholder="$t('ClickToFill')"
              theme="no-border"
            />
            <!-- Text -->
            <div v-else class="d-flex flex-items-center px-3 py-2">
              <span
                v-text="showAdminPass ? form.admin_password : _ap"
                class="mr-2"
              />
            </div>
            <GeneratePassword v-if="isEditMode" v-model="form.admin_password" />
            <CheckPass :password="form.admin_password" />
            <ShowPassBtn @click="showAdminPass = $event" />
            <ClipboardButton
              v-if="form.admin_username.length != 0"
              :copy="form.admin_password"
            />
          </div>
        </div>

        <!-- Extra -->
        <div class="form-row">
          <div class="d-flex flex-items-end flex-content-between">
            <label v-text="$t('EXTRA')" />
            <div class="d-flex flex-items-center">
              <ClipboardButton :copy="form.extra" />
            </div>
          </div>
          <div class="d-flex">
            <VTextArea
              v-model="form.extra"
              :sensitive="isEditMode || showExtra"
              :placeholder="$t('ClickToFill')"
              :disabled="!showExtra"
            />
          </div>
        </div>

        <!-- Save & Cancel -->
        <div class="d-flex m-3" v-if="isEditMode">
          <VButton
            class="flex-1"
            theme="text"
            :disabled="loading"
            @click="isEditMode = false"
          >
            {{ $t("Cancel") }}
          </VButton>
          <VButton class="flex-1" type="submit" :loading="loading">
            {{ $t("Save") }}
          </VButton>
        </div>
      </form>
    </PerfectScrollbar>
  </div>
</template>

<script>
import { mapState, mapActions } from "vuex";
import DetailMixin from "@/tools/mixins/detail";

export default {
  mixins: [DetailMixin],

  data() {
    return {
      isEditMode: false,
      showPass: false,
      showHostingPass: false,
      showAdminPass: false,
      showExtra: false,
    };
  },

  beforeRouteUpdate(to, from, next) {
    this.isEditMode = false;
    this.showPass = false;
    this.showHostingPass = false;
    this.showAdminPass = false;
    next();
  },

  methods: {
    ...mapActions("Servers", ["Delete", "Update"]),

    onClickDelete() {
      const onSuccess = async () => {
        await this.Delete(this.form.id);
        const index = this.ItemList.findIndex(
          (item) => item.id == this.form.id
        );
        if (index !== -1) {
          this.ItemList.splice(index, 1);
        }
        this.$router.push({ name: "Servers", params: { openFirst: true } });
      };

      this.$request(onSuccess, this.$waiters.Servers.Delete);
    },

    async onClickUpdate() {
      const onSuccess = async () => {
        await this.Update({ ...this.form });
        this.$router.push({ name: "Servers", params: { refresh: true } });
      };

      await this.$request(onSuccess, this.$waiters.Servers.Update);
      this.isEditMode = false;
    },
  },

  computed: {
    ...mapState("Servers", ["Detail", "ItemList"]),

    loading() {
      return this.$wait.is(this.$waiters.Servers.Update);
    },

    copyContent() {
      return [
        `Title: ${this.form.title}`,
        `IP: ${this.form.ip}`,
        `Username: ${this.form.username}`,
        `Password: ${this.form.password}`,
        `URL: ${this.form.url}`,
        `Hosting Username: ${this.form.hosting_username}`,
        `Hosting Password: ${this.form.hosting_password}`,
        `Admin Username: ${this.form.admin_username}`,
        `Admin Password: ${this.form.admin_password}`,
      ].join("\n");
    },

    getTitle() {
      return this.form.title || this.form.url;
    },

    _password() {
      if (this.form.password.length != 0) {
        return "•".repeat(this.form.password.length);
      } else {
        return "No Password";
      }
    },

    _hp() {
      if (this.form.hosting_password.length != 0) {
        return "•".repeat(this.form.hosting_password.length);
      } else {
        return "No Password";
      }
    },

    _ap() {
      if (this.form.admin_password.length != 0) {
        return "•".repeat(this.form.admin_password.length);
      } else {
        return "No Password";
      }
    },

    _au() {
      if (this.form.admin_username.length != 0) {
        return this.form.admin_username;
      } else {
        return "No Admin Username";
      }
    },

    _hu() {
      if (this.form.hosting_username.length != 0) {
        return this.form.hosting_username;
      } else {
        return "No Hosting Username";
      }
    },
  },
};
</script>
