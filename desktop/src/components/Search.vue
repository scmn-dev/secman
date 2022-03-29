<template>
  <div class="search">
    <div class="search-wrapper">
      <input
        type="text"
        :value="searchQuery"
        @input="onInputSearchQuery"
        :placeholder="$t('Search Here')"
      />
      <VIcon class="ml-2" name="search" size="16px" />
    </div>
  </div>
</template>

<script>
import { mapMutations, mapState } from "vuex";
export default {
  name: "Search",
  computed: mapState(["searchQuery"]),

  methods: {
    ...mapMutations(["onInputSearchQuery"]),
  },
};
</script>

<style lang="scss">
.search {
  width: 295px;
  text-align: center;

  &-wrapper {
    position: relative;
    width: 100%;
    border-right: 1px solid $color-gray-400;

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
}
</style>
