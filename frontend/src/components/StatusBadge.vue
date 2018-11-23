<template>
  <span v-bind:class="['status-badge', bgClass]">{{ text | capitalize }}</span>
</template>

<script>
export default {
  name: 'StatusBadge',
  props: {
    text: String,
  },
  computed: {
    bgClass() {
      if (!this.text) {
        return '';
      }
      switch (this.text.toLowerCase()) {
        case 'running':
          return 'status__on';
        case 'stopped':
          return 'status__off';
        case 'changing':
          return 'status__changing';
        default:
      }
      return 'status__default';
    },
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

.status__changing {
  color: #e54304;
  background: #ffddb0;
}

.status__default {
  color: #333;
  background: #ddd;
}
</style>
