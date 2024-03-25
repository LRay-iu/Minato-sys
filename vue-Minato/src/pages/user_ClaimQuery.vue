<template>
    <div>
        <PageTop />
        <el-container class="Maincontent">
            <ClaimSide />
            <el-container style="width: 1220px;"></el-container>
        </el-container>
        <PageBotton />
    </div>
</template>

<script lang='ts' setup name='userClaimQuery'>
import PageTop from "@/components/PageTop.vue"
import PageBotton from "@/components/PageBotton.vue"
import ClaimSide from "@/components/ClaimSide.vue"

</script>

<style scoped>
.Maincontent {
    height: 100vh;
}
</style>