<template>
  <div class="detail-page">
    <div class="detail-page-header">
      <!-- Avatar -->
      <div class="detail-page-header-avatar">
        <CompanyLogo :url="form.title" />
      </div>
      <!-- Summary -->
      <div class="detail-page-header-summary">
        <span v-text="$t('New Note')" class="url" />
        <span
          v-text="$t('Please fill all the necessary fields')"
          class="email"
        />
      </div>
    </div>
    <!-- Content -->
    <PerfectScrollbar class="detail-page-content">
      <form class="form" @submit.stop.prevent="onClickSave">
        <!-- Title -->
        <div class="form-row">
          <label v-text="$t('TITLE')" />
          <VFormText
            v-model="form.title"
            v-validate="'required'"
            name="Title"
            :placeholder="$t('ClickToFill')"
            theme="no-border"
          />
        </div>
        <!-- Note -->
        <div class="form-row">
          <div class="d-flex flex-items-end flex-content-between">
            <label v-text="$t('NOTE')" />
            <div class="d-flex flex-items-center">
              <ClipboardButton :copy="form.note" />
              <ShowPassBtn @click="showNote = $event" />
            </div>
          </div>
          <div class="d-flex">
            <VTextArea
              v-model="form.note"
              :sensitive="!showNote"
              :placeholder="$t('ClickToFill')"
              :disabled="showNote"
              name="Note"
            />
          </div>
        </div>

        <!-- Save & Cancel -->
        <div class="d-flex m-3">
          <VButton
            class="flex-1"
            theme="text"
            :disabled="loading"
            @click="$router.back()"
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
import { mapActions } from "vuex";

export default {
  data() {
    return {
      showNote: false,
      form: {
        title: "",
        note: "",
      },
    };
  },

  computed: {
    loading() {
      return this.$wait.is(this.$waiters.Notes.Create);
    },
  },

  methods: {
    ...mapActions("Notes", ["Create", "FetchAll"]),

    onClickSave() {
      this.$validator.validate().then(async (result) => {
        if (!result) return;

        const onSuccess = async () => {
          await this.Create({ ...this.form });
          this.FetchAll();
          this.$router.push({ name: "Notes", params: { refresh: true } });
        };

        this.$request(onSuccess, this.$waiters.Notes.Create);
      });
    },
  },
};
</script>
