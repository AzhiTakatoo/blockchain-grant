<template>
  <div class="app-container">
    <el-alert type="success" style="margin-bottom: 20px;" center show-icon>
      <p>学号: {{ wyuUserId }}</p>
        <p>姓名: {{ wyuUserName }}</p>
          <p>请确保您申请助学金提交的材料真实有效</p>
          <p>家庭人均年收入，综测分数可以写小数</p>
    </el-alert>
    <el-form ref="proofdata" v-loading="loading" :model="proofdata" label-width="110px">
      <el-form-item label="学号" prop="wyuUserId">
        {{wyuUserId}}
      </el-form-item>
      <!-- <el-form-item label="学号" prop="stipendId">
        <el-input v-model="proofdata.stipendId" placeholder="请输入学号" style="width:60%;margin-bottom:10px;"></el-input>
      </el-form-item> -->
      <el-form-item label="家庭人均年收入" prop="annualHouseholdIncome">
        <el-input placeholder="请输入具体收入" v-model="proofdata.annualHouseholdIncome" style="width:60%;margin-bottom:10px;"></el-input>
      </el-form-item>
      <el-form-item label="今年综测分数" prop="comprehensiveEvaluation">
        <el-input v-model="proofdata.comprehensiveEvaluation" placeholder="请输入具体分数" style="width:60%;margin-bottom:10px;"></el-input>
      </el-form-item>
      <el-form-item label="今年义工时长" prop="volunteerTime">
        <el-input v-model="proofdata.volunteerTime" placeholder="请输入具体时长" style="width:60%;margin-bottom:10px;"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm('proofdata')" style="width:30%;margin-bottom:10px;">提交</el-button>
        <el-button @click="resetForm('ruleForm')" style="width:27%;margin-bottom:30px;">重置</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
  import { mapGetters } from 'vuex'
  import { createProofMaterial } from '@/api/grantWork'
  const formT = () => {
    return {
      stipendId: '',
      annualHouseholdIncome: 0,
      comprehensiveEvaluation: 0,
      volunteerTime: 0,
      stipendScore: 0
    }
  }

  export default {
    name: 'AddText',
    data() {
      return {
        proofdata: formT(),
        loading: false
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
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            this.$confirm('是否确认无误，立即提交?', '提示', {
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'success'
            }).then(() => {

              // if (this.$store.state.grantUser.token !== this.proofdata.stipendId) {
              //   return Promise.reject('输入学号跟当前登录学号不匹配')
              // }
              this.loading = true
              createProofMaterial({
                // stipendId: this.proofdata.stipendId,
                stipendId: this.wyuUserId,
                annualHouseholdIncome: parseFloat(this.proofdata.annualHouseholdIncome),
                comprehensiveEvaluation: parseFloat(this.proofdata.comprehensiveEvaluation),
                volunteerTime: parseFloat(this.proofdata.volunteerTime)
              }).then(response => {
                this.loading = false
                if (response !== null) {
                  this.$message({
                    type: 'success',
                    message: '提交成功!'
                  })
                } else {
                  this.$message({
                    type: 'error',
                    message: '提交失败!'
                  })
                }
                this.proofdata = formT()
              }).catch(_ => {
                this.loading = false
              })
            }).catch(error => {

              this.loading = false
              this.$message({
                type: 'info',
                message: error || '提交失败'
              })
            })
          } else {
            return false
          }
        })
      },
      resetForm(formName) {
        this.proofdata = formT()
      },
      selectGet(wyuUserId) {
        this.ruleForm.proprietor = wyuUserId
      }
    }
  }
</script>

<style scoped>
</style>