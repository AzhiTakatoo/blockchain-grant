<template>
  <div class="container">
    <el-alert type="success" style="margin-bottom: 20px;" center show-icon>
      <p>学号: {{ wyuUserId }}</p>
        <p>姓名: {{ wyuUserName }}</p>
          <p>请确保您申请助学金提交的材料真实有效</p>
          <p>输入正确的权限密码，解锁查看获奖学金的名单</p>
    </el-alert>
    <div v-if="proofMaterialList.length==0" style="text-align: center;">
      <el-alert style="margin-bottom: 20px;" title="查询不到数据" type="warning" />
    </div>
    <div class="container1">
      <el-form v-if="(signPower !== '180806')||(roles[0] === 'admin')" ref="power" v-loading="loading" label-width="110px">
        <el-form-item label="权限" prop="power">
          <el-input placeholder="请输入权限密码" v-model="powerdata.power" style="width:60%;margin-bottom:10px;" show-password></el-input>
        </el-form-item>
        <el-button :loading="loading" type="primary" style="width:45%;margin-bottom:30px;" @click.native.prevent="unlock">解锁</el-button>
      </el-form>
    </div>

    <el-row v-if="signPower === '180806'" v-loading="loading" :gutter="20">
      <el-col v-for="(val,index) in proofMaterialList" :key="index" :span="6" :offset="1">
        <el-card class="rank-card">
          <div slot="header" class="clearfix">
            学号:
            <span style="color: rgb(255, 0, 0);">{{ val.stipendId }}</span>
          </div>

          <div class="item">
            <el-tag>家庭人均年收入: </el-tag>
            <span>{{ val.annualHouseholdIncome }} 元</span>
          </div>
          <div class="item">
            <el-tag type="success">综合测评: </el-tag>
            <span>{{ val.comprehensiveEvaluation }} 分</span>
          </div>
          <div class="item">
            <el-tag type="warning">义工时长: </el-tag>
            <span>{{ val.volunteerTime }} 小时</span>
          </div>
          <div class="item">
            <el-tag type="danger">助学金评定得分: </el-tag>
            <span>{{ val.stipendScore }} 分</span>
          </div>
            <el-button type="text" @click="queryPhoto(val.stipendId)">查看详情</el-button>
            <el-divider direction="vertical" />
            <el-button type="text" @click="open">提出质疑</el-button>
        </el-card>
      </el-col>
    </el-row>
  </div>

</template>

<script>
  import { mapGetters } from 'vuex'
  import { createQueryStipendRanking } from '@/api/rank'
  import { queryProofMaterial, queryPhotoMaterial } from '@/api/grantWork'
  import { updatePower, setPower } from '@/api/vote'
  const formT = () => {
    return {
      stipendId: '',
      stipendScore: 0,
      ranking: 0,
      grade: '',
    }
  }
  export default {
    name: 'Rank',
    data() {
      return {
        loading: false,
        loadingDialog: false,
        stipendrankingdata: formT(),
        proofMaterialList: [],
        photodata: {
          photo: '',
          wyuUserId: ''
        },
        powerdata: {
          power: ""
        },
        signPower: "",
      }
    },
    computed: {
      ...mapGetters([
        'wyuUserId',
        'roles',
        'wyuUserName',
      ])
    },
    created() {
      setPower().then(response => {
          this.signPower = response
        }).catch(() => {
          this.loading = false
        }),
        queryProofMaterial().then(response => {
          if (response !== null) {
            this.proofMaterialList = response
            console.log(this.proofMaterialList);

          }
          this.loading = false
        }).catch(_ => {
          this.loading = false
        })
    },
    methods: {
      unlock() {
        if (this.powerdata.power) {
          this.loading = true
          updatePower({ Power: this.powerdata.power }).then(response => {
              this.signPower = response
              this.loading = false
            }).catch(() => {
              this.loading = false
            }),
            queryProofMaterial().then(response => {
              this.proofMaterialList = response
            }).catch(() => {
              this.loading = false
            })
        } else {
          this.$message('请输入正确权限密码')
        }
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
        this.$refs[formName].resetFields()
      },
      open() {
        this.$alert('email:zhipingZeng_43@163.com', '监督部门联系方式', {
          confirmButtonText: '确定',
          callback: action => {
            this.$message({
              type: 'info',
              message: `action: ${ action }`
            });
          }
        });
      }
    }
  }
</script>

<style>
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

  .rank-card {
    width: 230px;
    height: 340px;
    margin: 18px;
  }
</style>