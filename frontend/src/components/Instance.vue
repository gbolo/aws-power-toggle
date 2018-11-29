<template>
  <div class="instance">
    <div class="container">
      <span>{{instance.name}}</span>
      <ToggleSwitch v-bind:isChecked="isOn" @click.native="toggleInstance(instance.id)" />
    </div>
    <div class="container">
      <span class="instance__type">{{instance.instance_type}}</span>
      <span>
        <clr-icon shape="cpu" size="24"></clr-icon> {{instance.vcpu}}
        <clr-icon shape="memory" size="24"></clr-icon> {{instance.memory_gb}}
      </span>
    </div>
  </div>
</template>

<script>
import InstancesApi from '@/services/api/Instances';
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
  },
  methods: {
    toggleInstance(id) {
      if (this.isOn) {
        InstancesApi.stopInstance(id);
      } else {
        InstancesApi.startInstance(id);
      }
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
    align-items: center;
    .instance__type {
      color: #a4a4a4;
    }
  }
}
</style>
