<template>
  <div class="environment-list__container">
    <Environment v-for="env in filteredEnvs" v-bind:env="env" :key="env.Name"/>
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
    filters: Array,
  },
  computed: {
    filteredEnvs() {
      return this.environments.filter(
        env => this.filters.some(
          (f) => {
            if (typeof (f.value) === 'string') {
              return (env[f.field] || '').toLowerCase() === f.value.toLowerCase();
            }
            return env[f.field] === f.value;
          },
        ),
      );
    },
  },
};

</script>

<style scoped lang="scss">
.environment-list__container {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  justify-content: center;
}
</style>
