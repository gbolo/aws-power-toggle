<template>
  <span v-bind:class="[isRunning ? 'status__on' : 'status__off', 'status-badge']">
    {{ text | capitalize }}
  </span>
</template>

<script>
export default {
  name: 'StatusBadge',
  props: {
    text: String,
    isRunning: Boolean,
  },
};
</script>

<style scoped lang="scss">
.status-badge {
  color: white;
  padding: 4px 8px;
  border-radius: 8px;
}
.status__on {
  color: #417505;
  background: #dcffb8;
}
.status__off {
  color: #ef0078;
  background: #fbe2f0;
}
</style>
