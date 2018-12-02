<template>
  <div
    v-bind:class="[ message.length ? 'show' : '']"
    id="snackbar"
  >
    <span class="message">{{message}}</span>
    <button
      class="x-btn"
      @click="clearError"
    >
      <clr-icon
        shape="window-close"
        size="22"
      ></clr-icon>
    </button>
  </div>
</template>

<script>
export default {
  name: 'Snackbar',
  props: {
    message: String,
  },
  methods: {
    clearError() {
      this.$store.dispatch('clearError');
    },
  },
};
</script>

<style lang="scss" scoped>
#snackbar {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  visibility: hidden;
  width: 420px;
  max-height: 210px;
  margin-left: -210px;
  background-color: #fbe2f0;
  color: #ef0078;
  text-align: center;
  border-radius: 8px;
  padding: 16px;
  position: fixed;
  z-index: 999;
  left: 50%;
  bottom: 16px;

  @media screen and (max-width: 450px) {
    width: 100%;
    margin-left: -50%;
    border-radius: 0;
  }

  .message {
    margin: auto 0;
  }

  .x-btn {
    cursor: pointer;
    color: #ef0078;
    margin: auto 0;
    padding: 16px;
    background: none;
    outline-style: none;
    border: none;

    &:hover {
      color: #880061;
    }
  }
}

#snackbar.show {
  visibility: visible;
  animation: fadein 0.5s;
}

@keyframes fadein {
  from {
    bottom: 0;
    opacity: 0;
  }
  to {
    bottom: 16px;
    opacity: 1;
  }
}
</style>
