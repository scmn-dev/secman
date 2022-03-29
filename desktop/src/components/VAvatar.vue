<template>
  <div class="avatar">
    <div class="avatar-content" v-text="firstLettersOfName" />
  </div>
</template>

<script>
export default {
  name: "VAvatar",
  props: {
    name: String,
  },
  computed: {
    firstLettersOfName() {
      const [firstName] = this.name.split(" ");
      return `${firstName[0]}`;
    },
  },
};
</script>

<style lang="scss">
.avatar {
  width: 40px;
  height: 40px;
  min-width: 40px;
  min-height: 40px;
  border-radius: 10px;
  cursor: pointer;

  background-origin: border-box;
  background-clip: content-box, border-box;
  border: double 2px transparent;
  background-image: linear-gradient(black, black),
    linear-gradient(45deg, $color-secondary, $color-primary);

  &-content {
    display: flex;
    width: 100%;
    height: 100%;
    justify-content: center;
    align-items: center;
    border-radius: 6px;
    font-weight: 700;
    color: #fff;
  }
}

.avatar:hover,
.avatar:focus {
  background-image: linear-gradient(black, black),
    linear-gradient(45deg, $color-primary, $color-secondary);
}
</style>
