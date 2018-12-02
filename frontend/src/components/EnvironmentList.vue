<template>
  <div>
    <ul class="environment-list__filter-list">
      <li
        v-for="(fa, index) in itemFilters"
        :key="index"
      >
        <a
          v-bind:class="[fa.selected ? 'selected' : '']"
          @click="toggleFilter(index)"
        >{{fa.value}}</a>
      </li>
    </ul>
    <div class="environment-list__container">
      <Environment
        v-for="env in filteredEnvs"
        v-bind:env="env"
        :key="env.Name"
      />
    </div>
  </div>
</template>

<script>
import Environment from '@/components/Environment.vue';

export default {
  name: 'EnvironmentList',
  components: {
    Environment,
  },
  props: {
    environments: Array,
  },
  data() {
    return {
      itemFilters: [
        { field: 'state', value: 'Running', selected: true },
        { field: 'state', value: 'Stopped', selected: true },
        { field: 'state', value: 'Changing', selected: true },
        { field: 'state', value: 'Mixed', selected: true },
      ],
    };
  },
  computed: {
    enabledFilters() {
      return this.itemFilters.filter(f => f.selected);
    },
    filteredEnvs() {
      return this.environments.filter(env => this.enabledFilters.some((f) => {
        if (typeof f.value === 'string') {
          return (env[f.field] || '').toLowerCase() === f.value.toLowerCase();
        }
        return env[f.field] === f.value;
      }));
    },
  },
  methods: {
    toggleFilter(index) {
      this.itemFilters[index].selected = !this.itemFilters[index].selected;
    },
  },
};
</script>

<style scoped lang="scss">
.environment-list__filter-list {
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
      line-height: inherit;
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

.environment-list__container {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  justify-content: center;
}
</style>
