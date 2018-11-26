<template>
  <div class="env">

    <div class="env__header">
      <span class="env__name">{{ env.name }}</span>
      <StatusBadge v-bind:text="env.state" />
    </div>

    <div class="env__details-container">
      <div class="env__details">
        <font-awesome-icon class="icon" v-bind:icon="getProviderIcon" />
        <span>{{env.region}}</span>
      </div>
      <div class="env__details">
        <font-awesome-icon class="icon" icon="memory" />
        <span>{{env.total_memory_gb}} GB</span>
      </div>
      <div class="env__details">
        <font-awesome-icon class="icon" icon="microchip" />
        <span>{{env.total_vcpu}} cores</span>
      </div>
      <div class="env__details">
        <font-awesome-icon class="icon" icon="server" />
        <span>{{env.running_instances}}/{{env.total_instances}}</span>
      </div>

      <font-awesome-icon @click="toggleInstanceList" class="chevron" icon="angle-double-down" />

      <InstanceList v-if="showInstances" v-bind:instances="env.instances" />
    </div>

    <button v-if="!isRunning &&
        !isLoading" class="button start" @click="start(env.id)">
      <font-awesome-icon icon="play" />
    </button>
    <button v-if="isRunning && !isLoading" class="button stop" @click="stop(env.id)">
      <font-awesome-icon icon="stop" />
    </button>
    <button v-if="isLoading" class="button disabled">
      <font-awesome-icon icon="spinner" />
    </button>

    <div v-if="error" class="env__error-container">
      <p class="error-message">{{ error }}</p>
      <p class="clear-error-message" @click="clearError">Clear Error</p>
    </div>

  </div>
</template>

<script>
import EnvironmentsApi from '@/services/api/Environments';
import StatusBadge from '@/components/StatusBadge.vue';
import InstanceList from '@/components/InstanceList.vue';

export default {
  name: 'Environment',
  components: {
    StatusBadge,
    InstanceList,
  },
  data() {
    return {
      isLoading: false,
      error: '',
      showInstances: false,
    };
  },
  props: {
    env: Object,
  },
  computed: {
    isRunning() {
      return this.env.running_instances > 0;
    },
    getProviderIcon() {
      if (!this.env.provider) {
        '';
      }

      switch(this.env.provider.toLowerCase()) {
        case 'aws':
          return ['fab', 'aws'];
          default:
      }
      return '';
    }
  },
  methods: {
    clearError() {
      this.error = '';
    },
    toggleInstanceList() {
      this.showInstances = !this.showInstances;
    },
    start(id) {
      EnvironmentsApi.startEnvironment(id)
        .then((response) => {
          EnvironmentsApi.getEnvironmentDetails(id).then((response) => {
            this.env = response;
          }).catch((e) => {
            this.error = e.response.data.error;
          });
        })
        .catch((e) => {
          this.error = e.response.data.error;
        })
        .finally(() => {
          this.isLoading = false;
        });
    },
    stop(id) {
      EnvironmentsApi.stopEnvironment(id)
        .then((response) => {
          EnvironmentsApi.getEnvironmentDetails(id).then((response) => {
            this.env = response;
          }).catch((e) => {
            this.error = e.response.data.error;
          });
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
.env {
  min-width: 300px;
  background: white;
  border-radius: 20px;
  box-shadow: 0 10px 40px -10px rgba(0, 0, 0, 0.2);
  padding: 16px;
  margin: 16px;

  .env__header {
    padding-bottom: 16px;
    border-bottom: 1px solid #ddd;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    .env__name {
      margin: auto 0;
      font-weight: bold;
    }
  }

  .chevron {
    cursor: pointer;
  }
}

.env__details-container {
  padding: 16px 0;

  .env__details {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    margin: 4px 0;

    .icon {
      width: 20px;
    }
  }
}

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
  background: #09af00;

  &:hover {
    background: #008b00;
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

.env__error-container {
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
