<template>
  <div class='environment'>
    <div class='environment__header'>
      <span class='environment__name'>{{ env.Name }}</span>
      <StatusBadge v-bind:isRunning='isRunning' v-bind:text='env.State'/>
    </div>
    <div class='environment__content'>
      <p>Instances Running: {{env.RunningInstances}}/{{env.TotalInstances}}</p>
      <button
        v-if='!isRunning &&
        !isLoading'
        class='button start'
        @click='start(env.Name)'
      >Start</button>
      <button
        v-if='isRunning && !isLoading'
        class='button stop'
        @click='stop(env.Name)'
      >Stop</button>
      <button v-if='isLoading' class='button disabled'>...</button>
    </div>
    <div v-if='error' class='environment__error-container'>
      <p class='error-message'>{{ error }}</p>
      <p class='clear-error-message' @click='clearError'>Clear Error</p>
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
  data() {
    return {
      isLoading: false,
      error: '',
    };
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
    clearError() {
      this.error = '';
    },
    start(envName) {
      EnvironmentsApi.startEnvironment(envName)
        .then((response) => {
          console.log(response);
        })
        .catch((e) => {
          this.error = e.data.error;
        })
        .finally(() => {
          this.isLoading = false;
        });
    },
    stop(envName) {
      EnvironmentsApi.stopEnvironment(envName)
        .then((response) => {
          console.log(response);
        })
        .catch((e) => {
          this.error = e.response.data.error;
        })
        .finally(() => {
          this.isLoading = false;
        });
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
    border: none;
    outline-style: none;
    transition: all 0.4s cubic-bezier(0.2, 0.2, 0.2, 1.2);
  }

  .start {
    background: #90ee02;

    &:hover {
      background: #61d800;
    }
  }

  .stop {
    background: #ee0290;

    &:hover {
      background: #dd0074;
    }
  }

  .disabled {
    background: #ddd;
    cursor: wait;
  }
}

.environment__error-container {
  .error-message {
    color: #ef0078;
    max-width: 250px;
    text-align: center;
    margin: 16px auto 0;
  }
  .clear-error-message {
    max-width: 250px;
    color: #2d2d2d;
    cursor: pointer;
    border-bottom: 1px solid #ddd;
    display: inline-block;
    margin: 8px auto 0;

    &:hover {
      border-bottom: 1px solid #2d2d2d;
    }
  }
}
</style>
