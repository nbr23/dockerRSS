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
  <div class="bg-blue-100 flex h-screen w-full items-center justify-center ">
    <div class="px-16 py-10 max-sm:px-8">
      <Header></Header>
      <div class="container mx-auto justify-center flex px-5 py-1 md:flex-row mx-auto w-screen items-center">
        <main>
          <form>
            <label for="imageName" class="mt-2 text-lg text-gray-600">Image name:</label>
            <input v-model="imageName" type="text" name="imageName" placeholder="nbr23/dockerrss[:latest]"
              class=" w-full bg-white-100 rounded border bg-opacity-50 border-gray-300 focus:border-indigo-500 focus:bg-white focus:ring-2 focus:ring-blue-200 text-base outline-none text-gray-700 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out">
            <label for="feed-url" class="mt-2 text-lg text-gray-600">Feed URL:</label>
            <input type="text" id="feed-url" name="feed-url" disabled :value=getRssFeedURL()
              class="w-full bg-white-100 rounded border bg-opacity-50 border-gray-300 focus:bg-transparent focus:border-indigo-500 text-base outline-none text-gray-700 py-1 px-3 leading-8">

            <div class="mt-1 flex justify-center text-lg text-black">
              <button v-on:click.prevent=copyUrlToClipboard(getRssFeedURL())
                class="inline-flex text-white bg-indigo-500 border-0 py-2 px-6 focus:outline-none hover:bg-indigo-600 rounded text-lg">Copy</button>
            </div>
          </form>
        </main>
      </div>
    </div>
  </div>
</template>

<style scoped>
header {
  line-height: 1.5;
}
</style>
