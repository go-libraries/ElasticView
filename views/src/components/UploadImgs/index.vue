<template>

  <div id="upload">
    <!--elementui的上传图片的upload组件-->
    <el-upload
      ref="upload"
      class="upload-demo"
      :file-list="fileList"
      list-type="picture-card"
      :action="uploadUrl"
      :limit="limit"
      :on-success="UploadSuccess"
      :on-preview="handlePictureCardPreview"
      :on-exceed="exceedHandle"
      :auto-upload="true"
      :on-remove="removeHandle"
      :multiple="true"
    >
      <i class="el-icon-plus" />
    </el-upload>
    <!--展示选中图片的区域-->
    <el-dialog :visible.sync="dialogImgVisible" style="position: absolute;top: 0px;z-index: 25555">
      <img
        width="100%"
        :src="dialogImageUrl"
        alt=""
      >
    </el-dialog>
  </div>

</template>

<script>
export default {
  name: 'Index',
  props: [
    'limit',
    'data',
    'index'
  ],
  data() {
    return {
      dialogImageUrl: '',
      dialogImgVisible: false,
      fileList: this.data,
      param: '', // 表单要提交的参数
      src: '' // 展示图片的地址
    }
  },
  computed: {
    uploadUrl() {
      return process.env.VUE_APP_BASE_API + '/api/upload/img'
    }
  },
  methods: {
    // 3，点击文件列表中已上传的文件时的钩子
    handlePictureCardPreview(file) {
      this.$emit('handlePictureCardPreview', file)
    },
    UploadSuccess(response, file, fileList) {
      this.fileList = fileList
      console.log(this.fileList, 'data')
      this.$emit('changeFileList', this.fileList, this.index)
    },
    // 5设置超过9张图给与提示
    exceedHandle() {
      this.$message({
        type: 'error',
        message: '您现在选择已超过' + this.limit + '张图，请重新选择'
      })
    },
    removeHandle(file, fileList) {
      this.fileList = fileList
      this.$emit('changeFileList', this.fileList, this.index)
    }
  }
}
</script>

<style scoped>

</style>
