<script setup>
import Header from './components/Header.vue'
</script>

<script>
export default {
  data() {
    return {
      imageName: '',
      currentUrl: location.toString(),
    }
  },
  methods: {
    getRssFeedURL: function () {
      return this.currentUrl + "tags/" + this.imageName
    },
    copyUrlToClipboard: function (str) {
      navigator.clipboard.writeText(str);
    }
  }
}
</script>

<template>
  <Header></Header>
  <main>
    <p>Enter the name of the Docker image to monitor:</p>
    <form>
      <label for="imageName">Image name: </label>
      <input v-model="imageName" type="text" name="imageName" placeholder="nbr23/dockerrss[:latest]">
      <p v-if="imageName">
        Url: <input :value=getRssFeedURL() disabled /><br />
        <button v-on:click.prevent=copyUrlToClipboard(getRssFeedURL())>Copy</button>
      </p>
    </form>
  </main>
</template>

<style scoped>
header {
  line-height: 1.5;
}
</style>
