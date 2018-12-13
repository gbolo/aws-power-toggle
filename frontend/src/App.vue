<template>
  <div id="app">
    <Header v-bind:version="version" />
    <button
      class="refresh-btn"
      @click="refresh"
    >
      <clr-icon
        shape="sync"
        size="16"
      ></clr-icon>
    </button>

    <EnvironmentList v-bind:environments="environments" />
    <Snackbar v-bind:message="error" />
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
}
</style>
