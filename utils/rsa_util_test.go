package utils

import (
	"fmt"
	"testing"
)

func TestRSAUtil(t *testing.T) {
	// 示例用法
	data := []byte("data to be verified")

	RAS_PUBLIC_KEY := "MIGCfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCSMHaZS36vPhrOtp8EnmYEnAL5Mu1xrCqUIKYb3avA/cV018P7tEmPJ2DTI0LbSyi8wKFZev6Y4DBYfwDUDieqyLVRj+O1cvV12PTANeJb4/uHZFAOMbdZpNzPx32cFBB/Q78+67V4QZnByT9Mo4QHZlhGjs/ERcklmOrgumu+kQIDAQAB"
	RAS_PRIVATE_KEY := "MIIYCdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAJIwdplLfq8+Gs62nwSeZgScAvky7XGsKpQgphvdq8D9xXTXw/u0SY8nYNMjQttLKLzAoVl6/pjgMFh/ANQOJ6rItVGP47Vy9XXY9MA14lvj+4dkUA4xt1mk3M/HfZwUEH9Dvz7rtXhBmcHJP0yjhAdmWEaOz8RFySWY6uC6a76RAgMBAAECgYAvw7/sZFtXHL6bjdW1J6ADn4nlLDuiqXMcBPLhJfNpUkFC5QC26/gg2ufq9Jlyl0MPMQ1G9EXdY1rIf26g2qMgzAcnltKJ2dY39GOFbDtW5Dww6cdK9JtR3a5GQejsxSf2UqK9tyA3aZ1OtICWz6fa+Rw48+YA5pBuGsWt0cRhMQJBAOyl8Gp57UP9IpRUmCEhLoH5oCxaZ0ju0/xWtSXLooQ/SEbDW5XPE88kVBuWfbKhl7f3vcrwDTG85Km5wkYjeD0CQQCeJNT7Wk2+hGJ97xN8JbhnoB5bZ50VTT6q2nxUHV+lyE1luVovcL7IpfB/AY+oiOIVFUxmrrNaeRAzr8pzWPDlAkBxH2GthFtHBNpizY1rSNFSkGFg0lZNJt1u4oP1bUJitV13ditxkWuGuXb7ORUdLuG3r1WqjNXB0On9uC6GGK6BAkBclXObM+MQBrEiyTS/GdY71KHxIVf1gKOPoxnmpMu6YuntA/aoj3kiPwPtVxyjrn+tmCqCcwTNktLJb8E2hnuBAkAsDY69Yy4IPKI7baGtPL0SRGoDz8s7AtNc2dMQMysi1YXwyTUaGZmWa7nsZXU7mDYOIoe/+AwRhY6obhHrRRYe"

	ss, _ := SignSHA256RSA(data, RAS_PRIVATE_KEY)
	valid, err := VerifySHA256RSA(data, RAS_PUBLIC_KEY, ss)
	if err != nil {
		fmt.Printf("Error verifying signature: %v\n", err)
		return
	}

	fmt.Printf("Signature valid: %v\n", valid)
}
