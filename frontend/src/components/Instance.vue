<template>
  <div class="instance">
    <div class="container">
      <span>{{instance.name}}</span>
      <ToggleButton
        v-if="!isUnknownState"
        :disabled="isLoading"
        :value="isOn"
        :sync="true"
        :color="'#09af00'"
        :width="40"
        :height="20"
        @change="toggleInstance"
      />
    </div>
    <div class="container">
      <span class="instance__type">{{instance.instance_type}}</span>
      <span>
        <clr-icon shape="cpu" size="24"></clr-icon>
        {{instance.vcpu}}
        <clr-icon shape="memory" size="24"></clr-icon>
        {{instance.memory_gb}}
      </span>
    </div>
  </div>
</template>

<script>
import ToggleButton from 'vue-js-toggle-button/src/Button.vue';

export default {
  name: 'Instance',
  components: {
    ToggleButton,
  },
  props: {
    envId: String,
    instance: Object,
  },
  computed: {
    isUnknownState() {
      return (
        this.instance.state.toLowerCase() !== 'running'
        && this.instance.state.toLowerCase() !== 'stopped'
      );
    },
    isLoading() {
      return this.$store.getters.isInstanceLoading(this.instance.id);
    },
  },
  data() {
    return {
      isOn: this.instance.state.toLowerCase() === 'running',
    };
  },
  methods: {
    toggleInstance({ value }) {
      const { id } = this.instance;
      const { envId } = this;
      this.isOn = value;
      if (!value) {
        this.$store.dispatch('stopInstance', { id, envId });
      } else {
        this.$store.dispatch('startInstance', { id, envId });
      }
    },
  },
  watch: {
    instance() {
      this.isOn = this.instance.state.toLowerCase() === 'running';
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
