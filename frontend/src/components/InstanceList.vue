<template>
  <transition name="expand">
    <ul v-if="show" class="instance-list">
      <li v-for="(instance, index) in instances" :key="index">
        <Instance v-bind:instance="instance" v-bind:envId="envId"/>
      </li>
    </ul>
  </transition>
</template>

<script>
import Instance from '@/components/Instance.vue';

export default {
  name: 'InstanceList',
  components: {
    Instance,
  },
  props: {
    envId: String,
    instances: Array,
    show: Boolean,
  },
};
</script>

<style lang="scss" scoped>
.instance-list {
  max-height: 445px;
  overflow: auto;
  list-style: none;
  padding: 0 8px;

  li {
    opacity: 1;
    transition: opacity 0.75s;
  }

  & li:first-child {
    border-top: 1px solid #ddd;
  }
}

.expand-enter {
  max-height: 0;
  margin: 0;
}
.expand-enter-to {
  max-height: 445px;
  transition: max-height 0.75s;
  margin: 0 auto 16px auto;
}

.expand-leave {
  max-height: 445px;
  margin: 0 auto 16px auto;
}
.expand-leave-to {
  max-height: 0;
  transition: max-height 0.75s;
  margin: 0;
}
</style>
