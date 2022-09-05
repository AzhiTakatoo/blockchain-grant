<template>
  <div class="app-container">
    <el-alert type="success" style="margin-bottom: 20px;" center show-icon>
      <p>学号: {{ wyuUserId }}</p>
        <p>姓名: {{ wyuUserName }}</p>
          <p>请确认您申请助学金修改的材料真实有效</p>
          <p>左边为原来提交的数据</p>
    </el-alert>
    <div style="display:flex; ">
      <div style="flex: 0 0 250px">
        <el-card class="updatetext-card">
          <div slot="header" class="clearfix">
            学号:
            <span style="color: rgb(255, 0, 0);">{{ proofCardData[0].stipendId }}</span>
          </div>

          <div class="item">
            <el-tag>家庭人均年收入: </el-tag>
            <span>{{ proofCardData[0].annualHouseholdIncome }} 元</span>
          </div>
          <div class="item">
            <el-tag type="success">综合测评: </el-tag>
            <span>{{ proofCardData[0].comprehensiveEvaluation }} 分</span>
          </div>
          <div class="item">
            <el-tag type="warning">义工时长: </el-tag>
            <span>{{ proofCardData[0].volunteerTime }} 小时</span>
          </div>
          <div class="item">
            <el-tag type="danger">助学金评定得分: </el-tag>
            <span>{{ proofCardData[0].stipendScore }} 分</span>
          </div>
          <el-button type="text" @click="queryPhoto(proofCardData[0].stipendId)">查看详情</el-button>
        </el-card>
      </div>

      <div style="flex: 1; padding:18px">
        <el-form ref="proofdata" v-loading="loading" :model="proofdata" label-width="110px">
          <el-form-item label="学号" prop="wyuUserId">
            {{this.$store.state.grantUser.token}}
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

    </div>
  </div>
</template>

<script>
  import { mapGetters } from 'vuex'
  import { updateProofMaterial, queryProofMaterialOnly, queryPhotoMaterial } from '@/api/grantWork'
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
    name: 'UpdateProofMaterial',
    data() {
      return {
        proofdata: formT(),
        photodata: {
          photo: '',
          wyuUserId: ''
        },
        proofCardData: [],
        ruleForm: {
          proprietor: '',
          totalArea: 0,
          livingSpace: 0
        },
        loading: false
      }
    },
    computed: {
      ...mapGetters([
        'wyuUserId',
        'wyuUserName',
        'roles'
      ])
    },
    created() {
      queryProofMaterialOnly({ stipendId: this.wyuUserId }).then(response => {
        if (response !== null) {
          this.proofCardData = response

        }
        this.loading = false
      }).catch(_ => {
        this.loading = false
      })
    },
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
              updateProofMaterial({
                stipendId: this.$store.state.grantUser.token,
                // stipendId: this.proofdata.stipendId,
                annualHouseholdIncome: parseFloat(this.proofdata.annualHouseholdIncome),
                comprehensiveEvaluation: parseFloat(this.proofdata.comprehensiveEvaluation),
                volunteerTime: parseFloat(this.proofdata.volunteerTime)
              }).then(response => {
                this.loading = false
                if (response !== null) {
                  this.$message({
                    type: 'success',
                    message: '修改成功!'
                  })
                } else {
                  this.$message({
                    type: 'error',
                    message: '修改失败!'
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
                message: error || '修改失败'
              })
            })
          } else {
            return false
          }
        })
      },
      queryPhoto(photoId) {
        queryPhotoMaterial({ wyuUserId: photoId }).then(response => {
          this.photodata.wyuUserId = response[0].wyuUserId
          this.photodata.photo = response[0].photo
          // console.log(this.photodata.wyuUserId);
          const h = this.$createElement;
          this.$msgbox({
            title: '文件资料哈希值',
            message: h('p', null, [
              h('span', null, this.photodata.photo),
            ]),
          }).then(action => {});
        }).catch(() => {
          this.loading = false
        })
      },
      resetForm(formName) {
        this.proofdata = formT()
      },
    }
  }
</script>

<style scoped>
  .container {
    width: 100%;
    /* text-align: center; */
    min-height: 100%;
    overflow: hidden;
  }

  .container1 {
    width: 100%;
    text-align: center;
    min-height: 100%;
    overflow: hidden;
  }

  .tag {
    float: left;
  }

  .item {
    font-size: 14px;
    margin-bottom: 18px;
    color: #999;
  }

  .clearfix:before,
  .clearfix:after {
    display: table;
  }

  .clearfix:after {
    clear: both
  }

  .updatetext-card {
    width: 230px;
    height: 340px;
    margin: 18px;
  }
</style>