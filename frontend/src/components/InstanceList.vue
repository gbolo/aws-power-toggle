<template>
  <ul v-bind:class="['instance-list', show ? 'visible' : 'hidden']">
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
    show: Boolean,
  },
};
</script>

<style lang="scss" scoped>
.instance-list {
  overflow: auto;
  max-height: 0;
  list-style: none;
  transition: max-height 0.75s;
  padding: 0 8px;
  margin: 0;

  li {
    opacity: 0;
    transition: opacity 0.75s;
  }

  & li:first-child {
    border-top: 1px solid #ddd;
  }

  &.visible {
    margin: 0 auto 16px auto;
    max-height: 445px;
    overflow: auto;
    li {
      opacity: 1;
    }
  }
}
</style>
