<template>
  <div class="divBusca">
    <img src="../assets/search.png" id="btnBusca" alt="Buscar" />
    <input
      type="text"
      id="txtBusca"
      placeholder="search by tag"
      v-on:input="updateValue($event.target.value)"
    />
  </div>
</template>

<script>
export default {
  data() {
    return {};
  },
  methods: {
    onSearchSubmition(value) {
      this.$emit("searchSubmitted", value);
    }
  }
};
</script>

<style scoped>
.divBusca {
  border: solid 1px;
  border-radius: 15px;
  width: 300px;
  margin-left: 50px;
  margin-right: 50px;
  margin-top: 30px;
}

#txtBusca {
  float: right;
  background-color: transparent;
  padding-right: 5px;
  font-style: italic;
  font-size: 18px;
  border: none;
  height: 32px;
  width: 260px;
}

#btnBusca {
  width: 18px;
  margin-top: 10px;
  margin-bottom: 8px;
}
</style>
