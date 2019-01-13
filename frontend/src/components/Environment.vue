<template>
  <div class="env">
    <div class="env__header">
      <span class="env__name">{{ env.name }}</span>
      <StatusBadge v-bind:text="env.state"/>
    </div>

    <table class="env__details-table">
      <tr>
        <td>
          <div class="env__details">
            <clr-icon shape="cloud" size="24"></clr-icon>
            <span>{{env.region}}</span>
          </div>
        </td>
        <td>
          <div class="env__details">
            <clr-icon shape="cluster" size="24"></clr-icon>
            <span>{{env.running_instances}}/{{env.total_instances}}</span>
          </div>
        </td>
      </tr>
      <tr>
        <td>
          <div class="env__details">
            <clr-icon shape="cpu" size="24"></clr-icon>
            <span>{{env.total_vcpu}} cores</span>
          </div>
        </td>
        <td>
          <div class="env__details">
            <clr-icon shape="memory" size="24"></clr-icon>
            <span>{{env.total_memory_gb}} GB</span>
          </div>
        </td>
      </tr>
    </table>

    <clr-icon
      shape="angle-double"
      size="20"
      dir="down"
      @click="toggleInstanceList"
      v-bind:class="['chevron', this.showInstanceList ? 'rotate-m180': '']"
      icon="angle-double-down"
    />
    <InstanceList
      v-bind:show="showInstanceList"
      v-bind:instances="env.instances"
      v-bind:envId="env.id"
    />

    <div v-if="isLoading" class="btn-container">
      <clr-icon shape="hourglass" size="24"></clr-icon>
    </div>
    <template v-else>
      <div v-if="isChanging" class="btn-container">
        <clr-icon
          v-bind:class="['refresh-icon', shouldSpinIcon ? 'spin-icon' : '']"
          @click="refresh(env.id)"
          shape="sync"
          size="24"
        ></clr-icon>
      </div>
      <div v-else class="btn-container">
        <button
          v-if="!isRunning"
          v-bind:class="['button', 'start', isMixed ? 'mixed' : '']"
          @click="start(env.id)"
        >Start</button>
        <button
          v-if="!isStopped"
          v-bind:class="['button', 'stop', isMixed ? 'mixed' : '']"
          @click="stop(env.id)"
        >Stop</button>
      </div>
    </template>
  </div>
</template>

<script>
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
      showInstanceList: false,
      showInstanceListItems: false,
      shouldSpinIcon: false,
    };
  },
  props: {
    env: Object,
  },
  computed: {
    isRunning() {
      return this.$store.getters.isEnvironmentStateRunning(this.env.id);
    },
    isStopped() {
      return this.$store.getters.isEnvironmentStateStopped(this.env.id);
    },
    isMixed() {
      return this.$store.getters.isEnvironmentStateMixed(this.env.id);
    },
    isChanging() {
      return this.$store.getters.isEnvironmentStateChanging(this.env.id);
    },
    isLoading() {
      return this.$store.getters.isEnvironmentLoading(this.env.id);
    },
  },
  watch: {
    isLoading(current, previous) {
      if (!current && previous === true) {
        // artificial delay of 750ms, should be the same as
        // the duration of the spin keyframes animation in styles
        setTimeout(() => {
          this.shouldSpinIcon = false;
        }, 750);
      } else if (current && previous === false) {
        this.shouldSpinIcon = true;
      }
    },
  },
  methods: {
    toggleInstanceList() {
      this.showInstanceList = !this.showInstanceList;
    },
    start(id) {
      this.$store.dispatch('startEnvironment', id);
    },
    stop(id) {
      this.$store.dispatch('stopEnvironment', id);
    },
    refresh(id) {
      this.$store.dispatch('fetchEnvironmentDetails', id);
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
  align-self: flex-start;

  .refresh-icon {
    margin: auto;
    cursor: pointer;
  }

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

  .btn-container {
    min-height: 34px;
    display: flex;
    justify-content: center;
  }

  .chevron {
    cursor: pointer;
    display: block;
    margin: 16px auto;
    transition: all 0.4s cubic-bezier(0.2, 0.2, 0.2, 1.2);
  }
  .rotate-m180 {
    transform: rotate(-180deg);
  }

  @keyframes spin {
    from {
      transform: rotate(0deg);
    }
    to {
      transform: rotate(360deg);
    }
  }
  .spin-icon {
    animation: spin infinite 0.75s linear;
  }

  .env__details {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    margin: auto 0;
    align-items: center;

    .icon {
      margin: auto 0;
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
    transition: background 0.4s cubic-bezier(0.2, 0.2, 0.2, 1.2);
  }

  .start {
    background: #09af00;
    &:hover {
      background: #008b00;
    }
    &.mixed {
      border-radius: 24px 0 0 24px !important;
    }
  }

  .stop {
    background: #ee0290;
    &:hover {
      background: #dd0074;
    }
    &.mixed {
      border-radius: 0 24px 24px 0 !important;
    }
  }

  .disabled {
    background: #ddd;
    cursor: wait;
  }

  .env__details-table {
    width: 100%;
    border-collapse: collapse;

    td {
      border: 1px solid #ddd;
      padding: 8px;
      width: 50%;
    }
    tr:first-child td {
      border-top: 0;
    }
    tr td:first-child {
      border-left: 0;
    }
    tr td:last-child {
      border-right: 0;
    }
  }
}
</style>
