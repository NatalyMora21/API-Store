<template>
    <div class="hello">
    </div>
</template>

<script>
  import axios from "axios"
  export default {
    name: 'HelloWorld',
    data(){
      return 
    },
    async amounted() {
      try {
          var result = await axios({
            method: "POST",
            url: "http/localhost:3000"
          })
        
      } catch (error) {

        console.error(error)
        
      }
  
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
