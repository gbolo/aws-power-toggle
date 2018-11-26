<template>
  <span v-bind:class="['status-badge', bgClass]">{{ text | capitalize }}</span>
</template>

<script>
export default {
  name: 'StatusBadge',
  props: {
    text: String,
  },
  filters: {
    capitalize(value) {
      return value ? value.toString().charAt(0).toUpperCase() + value.slice(1) : '';
    },
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
        case 'mixed':
          return 'status__mixed';
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
  margin: auto 0;
  transition: all 0.4s cubic-bezier(0.2, 0.2, 0.2, 1.2);
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

.status__mixed {
  color: #4a26fd;
  background: #eee6ff;
}

.status__default {
  color: #333;
  background: #ddd;
}
</style>
