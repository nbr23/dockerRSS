<script setup>
import Header from './components/Header.vue'
</script>

<script>
export default {
  data() {
    return {
      imageName: '',
      currentUrl: location.toString(),
      imageTag: '',
      imagePlatform: '',
    }
  },
  methods: {
    getRssFeedURL: function (preview = false) {
      const queryArgs = [];
      if (this.imagePlatform != 'Any' && this.imagePlatform != '') {
        queryArgs.push('platform=' + this.imagePlatform);
      }
      if (preview) {
        queryArgs.push('format=plain');
      }
      return this.currentUrl + "tags/" + this.imageName + (this.imageTag == '' ? '' : ':' + this.imageTag) + (queryArgs.length > 0 ? ('?' + queryArgs.join('&')) : '');
    },
    copyUrlToClipboard: function (str) {
      navigator.clipboard.writeText(str);
    },
    previewUrl: function (str) {
      window.open(str, '_blank');
    }
  }
}
</script>

<template>
  <a href="https://github.com/nbr23/dockerRSS"><img width="149" height="149"
      src="/static/forkme_right_darkblue_121621.svg" class="attachment-full size-full absolute top-0 right-0"
      alt="Fork me on GitHub" data-recalc-dims="1"></a>
  <div class="bg-blue-100 flex h-screen w-full items-center justify-center">
    <div class="">
      <Header></Header>
      <div class="container mx-auto justify-center flex px-5 py-1 md:flex-row mx-auto w-screen items-center">
        <main>
          <form>
            <div class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4 flex flex-col my-2">
              <div class="-mx-3 md:flex mb-6">
                <div class="md:w-3/5 px-3 mb-6 md:mb-0">
                  <label class="block uppercase tracking-wide text-grey-darker text-xs font-bold mb-2" for="imageName">
                    Image Name (*):
                  </label>
                  <input v-model="imageName"
                    class="appearance-none block w-full bg-grey-lighter text-grey-darker border border-red rounded py-3 px-4 mb-3"
                    id="imageName" type="text" placeholder="nbr23/dockerrss">
                </div>
                <div class="md:w-1/5 px-3">
                  <label class="block uppercase tracking-wide text-grey-darker text-xs font-bold mb-2" for="imageTag">
                    Tag:
                  </label>
                  <input v-model="imageTag"
                    class="appearance-none block w-full bg-grey-lighter text-grey-darker border border-grey-lighter rounded py-3 px-4"
                    id="imageTag" type="text" placeholder="latest">
                </div>
                <div class="md:w-2/5 px-3">
                  <label class="block uppercase tracking-wide text-grey-darker text-xs font-bold mb-2"
                    for="imagePlatform">
                    Platform:
                  </label>
                  <input v-model="imagePlatform"
                    class="block appearance-none w-full bg-grey-lighter border border-grey-lighter text-grey-darker py-3 px-4 pr-8 rounded"
                    id="imagePlatform" placeholder="linux/arm64" />
                </div>
              </div>
            </div>

            <div class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4 flex flex-col my-2">
              <div class="-mx-3 md:flex mb-6">
                <div class="md:w-full px-3 mb-6 md:mb-0">
                  <label class="block uppercase tracking-wide text-grey-darker text-xs font-bold mb-2" for="imageName">
                    Feed URL:
                  </label>
                  <input disabled
                    class="appearance-none block w-full bg-gray-100 rounded border bg-opacity-100 border-gray-300 border rounded py-3 px-4 mb-3"
                    id="feed-url" name="feed-url" type="text" :value=getRssFeedURL()>
                  <div class="flex mx-auto justify-center mx-auto items-center">
                    <button v-on:click.prevent=previewUrl(getRssFeedURL(true))
                      class="text-white bg-indigo-300 border-0 py-2 px-10 mr-4 focus:outline-none hover:bg-indigo-400 rounded text-lg">Preview</button>
                    <button v-on:click.prevent=copyUrlToClipboard(getRssFeedURL())
                      class="text-white bg-indigo-500 border-0 py-2 px-10 ml-4 focus:outline-none hover:bg-indigo-600 rounded text-lg">Copy</button>
                  </div>
                </div>
              </div>
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
