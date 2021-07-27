package router

import (
	"github.com/gin-gonic/gin"
	. "dsp/src/dsp_resource/controller"

)

/**
 * @author  tianqiang
 * @date  2021/7/20 17:12
 */
func Routers(router *gin.RouterGroup) {
	//测试连通性接口
	resCluster := router.Group("res")
	{
		resCluster.POST("/telnet", TelnetRes)
		resCluster.POST("/resIsExists", ResIsExists)
		resCluster.POST("/resStoreIsExists", ResStoreIsExists)
		resCluster.POST("/resAccountIsExists", ResAccountIsExists)
	}
	/*
	  安全设备
	*/
	//secDeviceCluster := router.Group("sec")
	//{
	//	secDeviceCluster.POST("/list", ResSecDeviceList)
	//	secDeviceCluster.POST("/create", CreateResSecDevice)
	//	secDeviceCluster.POST("/update", UpdateResSecDevice)
	//	secDeviceCluster.POST("/delete", DeleteResSecDevice)
	//	/*	secDeviceCluster.POST("/export", ExportResSecDevice)
	//		secDeviceCluster.POST("/download", DownloadResSecDevice)
	//		secDeviceCluster.POST("/createMore", CreateDeviceMore)*/
	//}
	//
	//hostCluster := router.Group("host")
	//{
	//	hostCluster.POST("/list", ListHost)
	//	hostCluster.POST("/create", CreateHost)
	//	hostCluster.POST("/delete", DeleteHost)
	//	hostCluster.POST("/update", UpdateHost)
	//	//登录接口
	//	hostCluster.POST("/login", LoginHost)
	//	/*hostCluster.POST("/createMore", CreateHostMore)
	//	hostCluster.POST("/export",ExportHost)*/
	//}
	//storeCluster := router.Group("store")
	//{
	//	storeCluster.POST("/list", ListStore)
	//	storeCluster.POST("/create", CreateStore)
	//	storeCluster.POST("/delete", DeleteStore)
	//	storeCluster.POST("/update", UpdateStore)
	//	/*storeCluster.POST("/createMore", CreateStoreMore)
	//	storeCluster.POST("/export",ExportStore)*/
	//}
	//networkDeviceCluster := router.Group("network")
	//{
	//	networkDeviceCluster.POST("/list", ResNetworkDeviceList)
	//	networkDeviceCluster.POST("/create", CreateResNetworkDevice)
	//	networkDeviceCluster.POST("/delete", DeleteResNetworkDevice)
	//	networkDeviceCluster.POST("/update", UpdateResNetworkDevice)
	//	/*networkDeviceCluster.POST("/createMore", CreateNetworkMore)
	//	networkDeviceCluster.POST("/export",ExportDevice)*/
	//}
	///*
	//  资源账号管理
	//*/
	//
	//accountCluster := router.Group("account")
	//{
	//	accountCluster.POST("/list", ListAccount)
	//	accountCluster.POST("/create", CreateAccount)
	//	accountCluster.POST("/update", UpdateAccount)
	//	accountCluster.POST("/delete", DeleteAccount)
	//	accountCluster.POST("/resAcc/list", ListAccountByResId)
	//	/*accountCluster.POST("/export", ExportAccount)
	//	accountCluster.POST("/download", DownloadAccount)
	//	accountCluster.POST("/createMore", CreateAccountMore)*/
	//}
	///*
	//  资源账号和资源授权
	//*/
	///*	accountAuthorCluster := router.Group("accAuthor")
	//	{
	//		accountAuthorCluster.POST("/list", ListAccountAuthor)
	//		accountAuthorCluster.POST("/create", CreateAccountAuthor)
	//		accountAuthorCluster.POST("/update", UpdateAccountAuthor)
	//		accountAuthorCluster.POST("/delete", DeleteAccountAuthor)
	//		accountAuthorCluster.POST("/export", ExportAccountAuthor)
	//		accountAuthorCluster.POST("/download", DownloadAccountAuthor)
	//		accountAuthorCluster.POST("/createMore", CreateAccountAuthorMore)
	//	}*/
	///*
	//  多IP管理
	//*/
	///*resIpsCluster := router.Group("resIps")
	//{
	//	resIpsCluster.POST("/list", ListResIps)
	//	resIpsCluster.POST("/oper", CreateResIps)
	//	resIpsCluster.POST("/update", UpdateResIps)
	//	resIpsCluster.POST("/delete", DeleteResIps)
	//}*/
	///*
	//  业务系统信息管理
	//*/
	///*	businessCluster := router.Group("business")
	//	{
	//		businessCluster.POST("/list", ResBusinessList)
	//		businessCluster.POST("/create", CreateResBusiness)
	//		businessCluster.POST("/update", UpdateResBusiness)
	//		businessCluster.POST("/delete", DeleteResBusiness)
	//		businessCluster.POST("/export", ExportResBusiness)
	//		businessCluster.POST("/download", DownloadResBusiness)
	//		businessCluster.POST("/createMore", CreateBusinessMore)
	//	}*/
	//
	///*
	//  协议管理
	//*/
	//proCluster := router.Group("resPros")
	//{
	//	proCluster.POST("/list", ListHostProtocol)
	//	proCluster.POST("/oper", CreateHostProtocol)
	//}
}