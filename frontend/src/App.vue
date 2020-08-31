<template>
  <div id="app">
    <div class="appbody">
      <Header v-bind:version="version"/>
      <button class="refresh-btn" @click="refresh">
        <clr-icon
          v-bind:class="[shouldSpinIcon ? 'spin-icon' : '']"
          shape="sync"
          size="16">
        </clr-icon>
      </button>
      <EnvironmentList v-bind:environments="environments"/>
      <Snackbar v-bind:message="error"/>
      <v-dialog/>
    </div>
    <div class="footer"></div>
  </div>
</template>

<script>
import Header from '@/components/Header.vue';
import EnvironmentList from '@/components/EnvironmentList.vue';
import Snackbar from '@/components/Snackbar.vue';

export default {
  name: 'app',
  components: {
    EnvironmentList,
    Header,
    Snackbar,
  },
  data() {
    return {
      shouldSpinIcon: false,
    };
  },
  computed: {
    environments() {
      return this.$store.state.environments;
    },
    version() {
      return this.$store.state.version;
    },
    error() {
      return this.$store.state.error;
    },
    isLoading() {
      return this.$store.state.isLoading;
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
  created() {
    this.$store.dispatch('fetchAllEnvironmentsDetails');
    this.$store.dispatch('fetchVersion');
  },
  methods: {
    refresh() {
      this.$store.dispatch('fetchAllEnvironmentsDetails');
    },
  },
};
</script>

<style lang="scss">
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  background: #f5f4f9;
  height: 100vh;
  margin: 0 auto;
  width: 85%;
  max-width: 1000px;

  @media screen and (max-width: 450px) {
    width: 100%;
  }

  .error-message {
    color: #ef0078;
    max-width: 250px;
    text-align: center;
    margin: 16px auto 0;
  }

  .refresh-btn {
    cursor: pointer;
    font-size: 1em;
    padding: 8px 32px;
    border-radius: 24px;
    color: #fff;
    border: none;
    outline-style: none;
    transition: all 0.4s cubic-bezier(0.2, 0.2, 0.2, 1.2);
    background: #9965f4;
    line-height: 0;
    &:hover {
      background: #7e3ff2;
    }
  }

  .appbody {
    min-height: calc(100vh - 60px);
  }

  .footer {
    padding:20px 0px 0px 0px;
    height:60px;
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

  // Override vue-js-modal (<vdialog/>) css
  .v--modal-overlay {
    background: rgba(0, 0, 0, 0.65) !important;
    .vue-dialog {
      margin: auto;
      left: 0 !important;
      border-radius: 10px;

      @media screen and (max-width: 550px) {
        width: 100% !important;
        border-radius: 0px;
      }
    }
  }
}
.tooltip {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  display: block;
  z-index: 10000;
  .tooltip-inner {
    background: black;
    color: white;
    border-radius: 6px;
    padding: 5px 10px 4px;
  }
}
</style>
