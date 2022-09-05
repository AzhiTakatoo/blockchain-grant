<template>
  <div class="container">
    <el-alert type="success" style="margin-bottom: 20px;" center show-icon>
      <p>学号: {{ wyuUserId }}</p>
        <p>姓名: {{ wyuUserName }}</p>
          <p>在助学金获得者名单公示</p>
          <p>依据受助学生是否发挥助学金的作用，真实客观地给助学金获得者打分</p>
    </el-alert>
    <div v-if="stipendRankingList.length==0" style="text-align: center;">
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
      <el-col v-for="(val,index) in stipendRankingList" :key="index" :span="6" :offset="1">
        <el-card class="award-card">
          <div slot="header" class="clearfix">
            学号:
            <span style="color: rgb(255, 0, 0);">{{ val.stipendId }}</span>
          </div>

          <div class="item">
            <el-tag type="warning">排名: </el-tag>
            <span>{{ val.ranking }} </span>
          </div>
          <div class="item">
            <el-tag>获助等级: </el-tag>
            <span>{{ val.grade }} </span>
          </div>
          <div class="item">
            <el-tag type="danger">获助金额: </el-tag>
            <span>{{ val.money }} 分</span>
          </div>
          <div class="item">
            <el-tag type="danger">助学金评定得分: </el-tag>
            <span>{{ val.stipendScore }} 分</span>
          </div>
          <div v-if="roles[0] !== 'admin'">
            <el-button type="text" @click="handleOpenVoteDialog(val.stipendId)">助后评分</el-button>
          </div>
          <el-dialog title="助后评分" :visible.sync="dialogVisible" width="30%">
            <el-rate v-model="rate" allow-half="true" show-text>
            </el-rate>
            <span slot="footer" class="dialog-footer">
              <el-button @click="dialogVisible = false" size="small">取 消</el-button>
              <el-button type="primary" @click="handlevote()" size="small" :loading="voteLoadging">确 定</el-button>
            </span>
          </el-dialog>
        </el-card>
      </el-col>
    </el-row>
  </div>

</template>

<script>
  import { mapGetters } from 'vuex'
  import { createQueryStipendRanking } from '@/api/rank'
  import { updatePower, setPower, createVote } from '@/api/vote'
  const formT = () => {
    return {
      stipendId: '',
      stipendScore: 0,
      ranking: 0,
      grade: '',
    }
  }
  export default {
    name: 'Assess',
    data() {
      return {
        dialogVisible: false,
        activeId: '',
        voteId: '',
        voteLoadging: false,
        rate: null,
        loading: false,
        loadingDialog: false,
        stipendrankingdata: formT(),
        stipendRankingList: [],
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
        createQueryStipendRanking().then(response => {
          this.stipendRankingList = response
        }).catch(() => {
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
            createQueryStipendRanking().then(response => {
              this.stipendRankingList = response
            }).catch(() => {
              this.loading = false
            })
        } else {
          this.$message('请输入正确权限密码')
        }
      },
      resetForm(formName) {
        this.$refs[formName].resetFields()
      },
      handlevote() {
        const data = {
          stipendId: this.activeId,
          voteId: this.voteId,
          vote: this.rate
        }
        this.voteLoadging = true
        createVote(data).then(response => {
          this.dialogVisible = false
          this.activeId = ''
          this.rate = null
          this.voteLoadging = false
          this.$message({
            type: 'success',
            message: '评分成功'
          })
        }).catch(() => {
          this.activeId = ''
          this.rate = null
          this.dialogVisible = false
          this.voteLoadging = false
        })
      },
      handleOpenVoteDialog(stipendId) {
        this.voteId = this.wyuUserId;
        this.dialogVisible = true;
        this.activeId = stipendId;
      },

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

  .award-card {
    width: 230px;
    height: 320px;
    margin: 18px;
  }
</style>