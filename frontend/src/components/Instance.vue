<template>
  <div class="instance">
    <div class="container">
      <span>{{instance.name}}</span>
      <ToggleSwitch v-bind:isChecked="isOn" />
    </div>
    <div class="container">
      <span>{{instance.instance_type}}</span>
      <span>
        <font-awesome-icon class="icon" icon="memory" /> {{instance.memory_gb}}
        <font-awesome-icon class="icon" icon="microchip" /> {{instance.vcpu}}</span>
    </div>
  </div>
</template>

<script>
import ToggleSwitch from '@/components/ToggleSwitch';

export default {
  name: 'Instance',
  components: {
    ToggleSwitch,
  },
  props: {
    instance: Object,
  },
  computed: {
    isOn() {
      return this.instance.state.toLowerCase() === 'running';
    }
  }
};
</script>

<style lang="scss" scoped>
.instance {
  padding: 8px 0;
  border-bottom: 1px solid #ddd;
  .container {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
  }
}
</style>
