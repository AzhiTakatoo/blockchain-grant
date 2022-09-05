<template>
  <div class="app-container">
    <el-alert style="margin-bottom: 20px;" title="上传申请认证材料" type="warning" description="请将申请助学金相关的贫困认证资料，综合测评分数表，义工时长证明等材料以图片形式存在一个word或excel文件中，尽量压缩图片大小" show-icon>
    </el-alert>
    <el-form v-loading="loading" label-width="110px">
      <el-form-item label="学号" prop="wyuUserId">
        {{wyuUserId}}
      </el-form-item>
      <el-form-item label="申请材料" prop="annualHouseholdIncome">
        <el-upload class="upload" action="/" :on-change="handleChange" :limit="1" name="uploadPhotoMaterial" :auto-upload="false">
          <el-button size="small" type="primary">选择文件</el-button>
          <div slot="tip" class="el-upload__tip">建议上传docx/xlsx文件</div>
        </el-upload>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm()" style="width:30%;margin-bottom:10px;" size="small" :loading="isLoading">提交</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
  import { mapGetters } from 'vuex'
  import { createProofMaterial, createPhotoMaterial } from '@/api/grantWork'

  export default {
    name: 'ProofMaterial',
    data() {
      return {
        loading: false,
        photoMaterial: null,
        isLoading: false
      }
    },
    computed: {
      ...mapGetters([
        'wyuUserId',
        'wyuUserName'
      ])
    },
    created() {},
    methods: {
      async submitForm(formName) {
        this.isLoading = true
        await createPhotoMaterial(this.photoMaterial)
        this.isLoading = false
        this.$message({
          type: 'success',
          message: '提交成功'
        })
      },
      handleChange(file, fileList) {
        const formData = new FormData()
        formData.append('wyuUserId', this.wyuUserId)
        formData.append('uploadPhotoMaterial', fileList[0].raw)
        this.photoMaterial = formData
      },

    }
  }
</script>

<style scoped>
</style>