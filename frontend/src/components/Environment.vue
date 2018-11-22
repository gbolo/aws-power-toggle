<template>
    <div class="environment">
      <div class="environment__header">
        <span class="environment__name">{{ env.Name }}</span>
        <StatusBadge
          v-bind:isRunning="isRunning"
          v-bind:text="env.State"
        />
      </div>
      <div class="environment__content">
        <p>Instances Running: {{env.RunningInstances}}/{{env.TotalInstances}}</p>
        <button v-if="!isRunning" class="button start"
          @click="start(env.Name)">Start</button>
        <button v-if="isRunning" class="button stop"
          @click="stop(env.Name)">Stop</button>
      </div>
    </div>
</template>

<script>
import EnvironmentsApi from '@/services/api/Environments';
import StatusBadge from '@/components/StatusBadge.vue';

export default {
  name: 'Environment',
  components: {
    StatusBadge,
  },
  props: {
    env: Object,
  },
  computed: {
    isRunning() {
      return this.env.RunningInstances > 0;
    },
  },
  methods: {
    start(envName) {
      EnvironmentsApi.startEnvironment(envName);
    },
    stop(envName) {
      EnvironmentsApi.stopEnvironment(envName);
    },
  },
};
</script>

<style scoped lang="scss">
.environment {
  min-width: 300px;
  background: white;
  border-radius: 20px;
  box-shadow: 0 10px 40px -10px rgba(0, 0, 0, 0.2);
  padding: 16px;
  margin: 16px;
}
.environment__header {
  padding-bottom: 16px;
  border-bottom: 1px solid #ddd;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}
.environment__name {
  margin: auto 0;
  font-weight: bold;
}
.environment__content {
  .button {
    cursor: pointer;
    font-size: 1em;
    padding: 8px 32px;
    border-radius: 24px;
    color: white;
    background: #90ee02;
    border: none;
    outline-style: none;
  }
  .start {
    background: #90ee02;
  }
  .stop {
    background: #ee0290;
  }
}
</style>
