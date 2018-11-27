<template>
  <label class="switch">
    <input type="checkbox" v-model="checked" @click.stop>
    <span class="slider round"></span>
  </label>
</template>

<script>
export default {
  name: 'ToggleSwitch',
  props: {
    isChecked: Boolean,
  },
  data() {
    return {
      checked: this.isChecked,
    };
  },
};
</script>

<style scoped lang="scss">
.switch {
  position: relative;
  display: inline-block;
  width: 30px;
  height: 16px;
}

.switch input {
  opacity: 0;
  width: 0;
  height: 0;
}

.slider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #ccc;
  -webkit-transition: 0.4s;
  transition: 0.4s;
}

.slider:before {
  position: absolute;
  content: '';
  height: 12px;
  width: 12px;
  left: 2px;
  bottom: 2px;
  background-color: white;
  -webkit-transition: 0.4s;
  transition: 0.4s;
}

input:checked + .slider {
  background-color: #09af00;
}

input:focus + .slider {
  box-shadow: 0 0 1px #09af00;
}

input:checked + .slider:before {
  transform: translateX(12px);
}

.slider.round {
  border-radius: 34px;
}

.slider.round:before {
  border-radius: 50%;
}
</style>
