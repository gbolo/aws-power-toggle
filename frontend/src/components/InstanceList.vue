<template>
  <ul class="instance-list">
    <li v-for="(instance, index) in instances" :key="index">
      <Instance v-bind:instance="instance" />
    </li>
  </ul>
</template>

<script>
import Instance from '@/components/Instance.vue';

export default {
  name: 'InstanceList',
  components: {
    Instance,
  },
  props: {
    instances: Array,
  },
};
</script>

<style lang="scss" scoped>
.instance-list {
  list-style: none;
  padding: 0 8px;
  margin: 0 auto;
}
</style>
