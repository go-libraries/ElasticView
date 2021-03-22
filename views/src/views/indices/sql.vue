<template>
    <div>
      <sql-editor
        v-model="sql"
        styles="width: 100%"
      />
    </div>
</template>

<script>

    export default {
        name: "sql",
        data(){
          return {
              sql:"11111"
          }
        },
        components: {
            'SqlEditor': () => import('@/components/SqlEditor/index')
        },
        created() {

        }
    }
</script>

<style scoped>

</style>
