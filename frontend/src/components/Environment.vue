<template>
  <div class="env">

    <div class="env__header">
      <span class="env__name">{{ env.name }}</span>
      <StatusBadge v-bind:text="env.state" />
    </div>

    <table class="env__details-table">
      <tr>
        <td>
          <div class="env__details">
            <clr-icon
              shape="cloud"
              size="24"
            ></clr-icon>
            <span>{{env.region}}</span>
          </div>
        </td>
        <td>
          <div class="env__details">
            <clr-icon
              shape="cluster"
              size="24"
            ></clr-icon>
            <span>{{env.running_instances}}/{{env.total_instances}}</span>
          </div>
        </td>
      </tr>
      <tr>
        <td>
          <div class="env__details">
            <clr-icon
              shape="cpu"
              size="24"
            ></clr-icon>
            <span>{{env.total_vcpu}} cores</span>
          </div>
        </td>
        <td>
          <div class="env__details">
            <clr-icon
              shape="memory"
              size="24"
            ></clr-icon>
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
    />

    <button
      v-if="!isRunning"
      class="button start"
      @click="start(env.id)"
    >
      Start
    </button>
    <button
      v-if="isRunning"
      class="button stop"
      @click="stop(env.id)"
    >
      Stop
    </button>
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
    };
  },
  props: {
    env: Object,
  },
  computed: {
    isRunning() {
      return this.env.running_instances > 0;
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
    display: block;
    margin: 16px auto;
    transition: all 0.4s cubic-bezier(0.2, 0.2, 0.2, 1.2);
  }
  .rotate-m180 {
    transform: rotate(-180deg);
  }
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
</style>
