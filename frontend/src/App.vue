<template>
  <div id="app">
    <Header />
    <EnvironmentList v-bind:environments="environments"/>
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
    };
  },
  created() {
    EnviromentsApi.getEnvironments()
      .then((data) => {
        this.environments = data;
      });
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
  // background-image: linear-gradient(0deg, #647dee 0%, #7f53ac 100%);
  background: #f5f4f9;
  height: 100vh;
  margin: 0 auto;
  width: 85%;
  max-width: 1000px;
}
</style>
