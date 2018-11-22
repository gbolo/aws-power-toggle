<template>
  <div id="app">
    <Header />
    <EnvironmentList v-bind:environments="environments"/>
    <p v-if="error" class="error-message">{{ error }}</p>
  </div>
</template>

<script>
import EnviromentsApi from '@/services/api/Environments';
import Header from '@/components/Header.vue';
import EnvironmentList from '@/components/EnvironmentList.vue';

export default {
  name: 'app',
  components: {
    EnvironmentList,
    Header,
  },
  data() {
    return {
      environments: [],
      isLoading: false,
      error: '',
    };
  },
  created() {
    EnviromentsApi.getEnvironments()
      .then((data) => {
        this.environments = data;
      })
      .catch((e) => { this.error = e.error; })
      .finally(() => { this.isLoading = false; });
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

  .error-message {
    color: #ef0078;
    max-width: 250px;
    text-align: center;
    margin: 16px auto 0;
  }
}
</style>
