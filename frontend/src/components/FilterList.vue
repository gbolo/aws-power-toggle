<template>
  <ul>
    <li v-for="(fa, index) in filtersAvailable" :key="index">
      <a
        v-bind:class="[fa.selected ? 'selected' : '']"
        @click="toggleSelected(index)"
      >{{fa.value}}</a>
    </li>
  </ul>
</template>

<script>
export default {
  name: 'FilterList',
  props: {
    filters: Array,
  },
  data() {
    return {
      filtersAvailable: [
        { field: 'State', value: 'Running', selected: true },
        { field: 'State', value: 'Stopped', selected: true },
        { field: 'State', value: 'Changing', selected: true },
      ],
    };
  },
  created() {
    this.$emit('updatedFilters', this.filtersAvailable.filter(f => f.selected));
  },
  methods: {
    toggleSelected(index) {
      this.filtersAvailable[index].selected = !this.filtersAvailable[index].selected;
      this.$emit('updatedFilters', this.filtersAvailable.filter(f => f.selected));
    },
  },
};
</script>

<style scoped lang="scss">
ul {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  justify-content: center;
  text-align: center;
  padding: 0;
  margin: 16px 0;
  list-style: none;

  li {
    text-align: center;
    display: inline-block;
    margin: auto 8px;
    padding: 0;
    line-height: 2;

    a {
      font-size: 13px;
      line-height: 27px;
      color: rgb(102, 102, 102);
      cursor: pointer;
      display: inline-block;
      font-weight: normal;
      padding: 0px 16px;
      border-radius: 100px;
      transition: all 0.2s ease 0s;
      background-color: #ddd;

      &.selected {
        background-color: #2773ff;
        color: white;
      }
    }
  }
}
</style>
