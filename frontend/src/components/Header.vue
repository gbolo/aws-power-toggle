<template>
  <header>
    <div class="brand-container">
      <span
        role="img"
        aria-label="AWS Power-Toggle logo"
      >🔌✨</span>
      <div class="divider"></div>
      <span>AWS Power-Toggle</span>
    </div>
    <nav>
      <span class="version-container">{{versionLabel}}</span>
      <a
        target="_blank"
        href="https://github.com/gbolo/aws-power-toggle"
      >GitHub</a>
    </nav>
  </header>
</template>

<script>
export default {
  name: 'Header',
  props: {
    version: String,
  },
  data() {
    return {
      refreshing: false,
    };
  },
  computed: {
    versionLabel() {
      if (!this.version) {
        return '';
      }
      return /^\d/.test(this.version) ? `v${this.version}` : this.version;
    },
  },
};
</script>

<style scoped lang="scss">
header {
  min-height: 10vh;
  background-color: transparent;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  flex-wrap: wrap;
  font-size: 1.5em;
  color: #2d2d2d;
  padding: 16px;

  .brand-container {
    display: flex;
    flex-direction: row;
    margin: auto 0;
    @media screen and (max-width: 630px) {
      margin: auto;
    }

    .divider {
      background-color: #2d2d2d;
      height: 28px;
      width: 1px;
      margin: 0px 10px;
    }
  }

  nav {
    align-self: flex-end;
    margin: auto 0;
    @media screen and (max-width: 630px) {
      display: none;
    }
    .refresh {
      cursor: pointer;
    }

    .version-container {
      font-size: 0.8em;
      margin: 0 16px;
    }

    a {
      font-size: 0.8em;
      text-decoration: none;
      padding-bottom: 1px;
      color: #2d2d2d;

      &:hover {
        color: #5d5d5d;
        border-bottom: 1px solid #ddd;
      }
    }
  }
}
</style>
