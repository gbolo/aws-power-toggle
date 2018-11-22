<template>
    <div class="environment-list__container">
        <Environment
          v-for="environment in environments"
          v-bind:env="environment"
          :key="environment.Name"
        />
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
