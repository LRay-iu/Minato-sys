<template>
    <div>
        <PageTop />
        <el-container class="Maincontent">
            <InsuranceSide />
            <el-container style="width: 1220px;"></el-container>
        </el-container>
        <PageBotton />
    </div>
</template>

<script lang='ts' setup name='userInsuranceQuery'>
import PageTop from "@/components/PageTop.vue"
import PageBotton from "@/components/PageBotton.vue"
import InsuranceSide from "@/components/InsuranceSide.vue"

</script>

<style scoped>
.Maincontent {
    height: 100vh;
}
</style>