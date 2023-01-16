/*!***********************************************************************
*@file aciext.h
*@brief aci��չ����
*@author yfw
*@date 2011-11-22
*@note
*
*��Ϊaci.h��oracle oci.h�Ķ�����һ�µģ����Բ���aci.h�������չ�Ķ���
*����aci��չ�Ķ�������ڴ��ļ������
************************************************************************/
#ifdef __cplusplus
extern "C" {
#endif

#ifndef ACI_EXT_H
#define ACI_EXT_H

/*=============================Attribute Types===============================*/
/* 
Note: All attributes are global.  New external attibutes should be added to the end
of the list. Before you add an attribute see if an existing one can be 
used for your handle. 

If you see any holes please use the holes first. 

*/
/*===========================================================================*/
#define ACI_ATTR_INSERT_ROW_ROWID   420 //!< rowid of inserted row
#define ACI_ATTR_FETCH_ROWIDS       421 //!< rowids of bulkinsert row

#endif /* ACI_EXT_H */

#ifdef __cplusplus
}
#endif /* __cplusplus */
