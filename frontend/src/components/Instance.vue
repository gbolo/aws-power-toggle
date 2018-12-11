<template>
  <div class="instance">
    <div class="container">
      <span>{{instance.name}}</span>
      <ToggleSwitch
        v-if="!isMixedState"
        v-bind:isChecked="isOn"
        @click.native="toggleInstance(instance.id, envId)"
      />
    </div>
    <div class="container">
      <span class="instance__type">{{instance.instance_type}}</span>
      <span>
        <clr-icon
          shape="cpu"
          size="24"
        ></clr-icon> {{instance.vcpu}}
        <clr-icon
          shape="memory"
          size="24"
        ></clr-icon> {{instance.memory_gb}}
      </span>
    </div>
  </div>
</template>

<script>
import ToggleSwitch from '@/components/ToggleSwitch.vue';

export default {
  name: 'Instance',
  components: {
    ToggleSwitch,
  },
  props: {
    envId: String,
    instance: Object,
  },
  computed: {
    isMixedState() {
      return (
        this.instance.state.toLowerCase() !== 'running'
        && this.instance.state.toLowerCase() !== 'stopped'
      );
    },
  },
  data() {
    return {
      isOn: this.instance.state.toLowerCase() === 'running',
    };
  },
  methods: {
    toggleInstance(id, envId) {
      if (this.isOn) {
        this.$store.dispatch('stopInstance', { id, envId });
      } else {
        this.$store.dispatch('startInstance', { id, envId });
      }
      this.isOn = !this.isOn;
    },
  },
};
</script>

<style lang="scss" scoped>
.hide {
  display: none;
}
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
