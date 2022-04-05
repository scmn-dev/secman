<template>
  <button
    type="button"
    class="showpass-btn ml-1 trsn"
    v-tooltip="
      $t(isShow ? (isNote ? 'Show' : 'Hide') : isNote ? 'Hide' : 'Show')
    "
    v-on="inputListeners"
  >
    <VIcon
      class="c-pointer"
      :name="isShow ? (isNote ? 'eye' : 'eye-off') : isNote ? 'eye-off' : 'eye'"
      size="14px"
    />
  </button>
</template>

<script>
export default {
  name: "ShowPassBtn",

  props: {
    isNote: Boolean,
  },

  data() {
    return { isShow: false };
  },

  computed: {
    inputListeners() {
      return Object.assign({}, this.$listeners, {
        click: (event) => {
          this.isShow = !this.isShow;
          return this.$emit("click", this.isShow);
        },
      });
    },
  },
};
</script>

<style lang="scss">
.showpass-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 24px;
  height: 24px;
  border-radius: 4px;
  background-color: $color-gray-500;
  color: $color-gray-300;
}

.showpass-btn:hover {
  color: $color-secondary;
}
</style>
