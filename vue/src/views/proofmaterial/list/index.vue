<template>
  <div class="container">
    <el-alert type="success" center show-icon>
      <p>学号: {{ wyuUserId }}</p>
        <p>姓名: {{ wyuUserName }}</p>
          <p>请确保您申请助学金提交的材料真实有效</p>
          <p>下面列表将持续更新全部申请助学金的名单</p>
    </el-alert>
    <div v-if="proofMaterialList.length==0" style="text-align: center;">
      <el-alert title="查询不到数据" type="warning" />
    </div>
    <el-row v-loading="loading" :gutter="20">
      <el-col v-for="(val,index) in proofMaterialList" :key="index" :span="6" :offset="1">
        <el-card class="proofCertify-card">
          <div slot="header" class="clearfix">
            学号:
            <span style="color: rgb(255, 0, 0);">{{ val.stipendId }}</span>
          </div>
          <div class="item">
            <el-tag type="danger">申请时间: </el-tag>
            <span>{{ val.registerTime }}</span>
          </div>
          <el-rate v-if="roles[0] === 'admin'" />
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
  import { mapGetters } from 'vuex'
  import { queryProofCertify } from '@/api/grantWork'

  export default {
    name: 'proofMaterialList',
    data() {
      var checkArea = (rule, value, callback) => {
        if (value <= 0) {
          callback(new Error('必须大于0'))
        } else {
          callback()
        }
      }
      return {
        loading: true,
        loadingDialog: false,
        proofMaterialList: [],
        realForm: {
          price: 0,
          salePeriod: 0
        },
        rules: {
          price: [
            { validator: checkArea, trigger: 'blur' }
          ],
          salePeriod: [
            { validator: checkArea, trigger: 'blur' }
          ]
        },
        valItem: {}
      }
    },
    computed: {
      ...mapGetters([
        'wyuUserId',
        'roles',
        'wyuUserName',
        'balance'
      ])
    },
    created() {
      if (this.roles[0] === 'admin') {
        queryProofCertify().then(response => {

          if (response !== null) {
            this.proofMaterialList = response
            console.log(this.proofMaterialList);

          }
          this.loading = false
        }).catch(_ => {
          this.loading = false
        })
      } else {
        console.log(this.proofMaterialList);
        queryProofCertify().then(response => {
          console.log(response);

          if (response !== null) {
            this.proofMaterialList = response
            console.log(this.proofMaterialList);
          }
          this.loading = false
        }).catch(_ => {
          this.loading = false
        })
      }
    },
    methods: {
     
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

  .proofCertify-card {
    width: 280px;
    height: 130px;
    margin: 18px;
  }
</style>