<template>
  <header>
    <div class="brand-container">
      <span role="img" aria-label="AWS Power-Toggle logo">🔌✨</span>
      <div class="divider">
      </div><span>AWS Power-Toggle</span></div>
  </header>
</template>

<script>
export default {
  name: 'Header',
};
</script>

<style scoped lang="scss">
header {
  height: 10vh;
  background-color: transparent;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  -webkit-box-pack: center;
  justify-content: center;
  font-size: 1.5em;
  color: #2d2d2d;
  padding: 16px 16px 0 16px;
}
.brand-container {
  display: flex;
  flex-direction: row;
  .divider {
    background-color: #2d2d2d;
    height: 28px;
    width: 1px;
    margin: 0px 10px;
  }
}
</style>
