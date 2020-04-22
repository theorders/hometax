package hometax

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/imroc/req"
	"github.com/theorders/aefire"
)

type CorpInfoResponse struct {
	XMLName    xml.Name `xml:"map"`
	Text       string   `xml:",chardata"`
	ID         string   `xml:"id,attr"`
	TxprDscmNo string   `xml:"txprDscmNo"`
	Map        []struct {
		Text               string `xml:",chardata"`
		ID                 string `xml:"id,attr"`
		TxprDscmNoClCd     string `xml:"txprDscmNoClCd"`
		TxprDscmDt         string `xml:"txprDscmDt"`
		AbftChrgTxhfOgzCd  string `xml:"abftChrgTxhfOgzCd"`
		Tin                string `xml:"tin"`
		AbftChrgDutsCd     string `xml:"abftChrgDutsCd"`
		TxprClsfCd         string `xml:"txprClsfCd"`
		TxprNm             string `xml:"txprNm"`
		TxprDclsCd         string `xml:"txprDclsCd"`
		TxprDscmNoDplcClCd string `xml:"txprDscmNoDplcClCd"`
		TxprMntgYn         string `xml:"txprMntgYn"`
		TxprStatCd         string `xml:"txprStatCd"`
		TxprDscmNoEncCntn  string `xml:"txprDscmNoEncCntn"`
		DetailMsg          string `xml:"detailMsg"`
		Msg                string `xml:"msg"`
		Map                struct {
			Text                   string `xml:",chardata"`
			ID                     string `xml:"id,attr"`
			TxprDscmNo             string `xml:"txprDscmNo"`
			BmanOfbDt              string `xml:"bmanOfbDt"`
			EtxivPkcYn             string `xml:"etxivPkcYn"`
			TxofOgzCd              string `xml:"txofOgzCd"`
			LgnUserClCd            string `xml:"lgnUserClCd"`
			TxprDscmNoCnfrYn       string `xml:"txprDscmNoCnfrYn"`
			HaboCl                 string `xml:"haboCl"`
			NtplBmanAthYn          string `xml:"ntplBmanAthYn"`
			CharId                 string `xml:"charId"`
			SmprYn                 string `xml:"smprYn"`
			PubcUserNo             string `xml:"pubcUserNo"`
			ThofOgzCd              string `xml:"thofOgzCd"`
			Tin                    string `xml:"tin"`
			RprsTin                string `xml:"rprsTin"`
			BmanUnitEngeTrerJnngYn string `xml:"bmanUnitEngeTrerJnngYn"`
			RprsResno              string `xml:"rprsResno"`
			TxprClsfCd             string `xml:"txprClsfCd"`
			RprsFnm                string `xml:"rprsFnm"`
			UserClsfCd             string `xml:"userClsfCd"`
			DprtUserYn             string `xml:"dprtUserYn"`
			TxpAgnYn               string `xml:"txpAgnYn"`
			CnvrTin                string `xml:"cnvrTin"`
			AfaTxprYn              string `xml:"afaTxprYn"`
			SsnAltPsbYn            string `xml:"ssnAltPsbYn"`
			MpbNo                  string `xml:"mpbNo"`
			LgnClientIp            string `xml:"lgnClientIp"`
			UserId                 string `xml:"userId"`
			ChrgDutsCd             string `xml:"chrgDutsCd"`
			DataMaagClCd           string `xml:"dataMaagClCd"`
			CrpBmanAthYn           string `xml:"crpBmanAthYn"`
			LgnCertCd              string `xml:"lgnCertCd"`
			TxaaYn                 string `xml:"txaaYn"`
			UserCertClCd           string `xml:"userCertClCd"`
			WhlPmtMnpt             string `xml:"whlPmtMnpt"`
			NtplAthYn              string `xml:"ntplAthYn"`
			UserNm                 string `xml:"userNm"`
		} `xml:"map"`
		Code                   string `xml:"code"`
		Result                 string `xml:"result"`
		RprsYmdgClCd           string `xml:"rprsYmdgClCd"`
		ItmNm                  string `xml:"itmNm"`
		PfbPsenClCd            string `xml:"pfbPsenClCd"`
		PgBsTypeCd             string `xml:"pgBsTypeCd"`
		VatTxtnTrgtYn          string `xml:"vatTxtnTrgtYn"`
		Lfamt                  string `xml:"lfamt"`
		LdAdr                  string `xml:"ldAdr"`
		HofcTxprStatCd         string `xml:"hofcTxprStatCd"`
		SggEnglNm              string `xml:"sggEnglNm"`
		HofcRprsResnoDec       string `xml:"hofcRprsResnoDec"`
		EmplCnt                string `xml:"emplCnt"`
		OnsCpamt               string `xml:"onsCpamt"`
		HoAdr                  string `xml:"hoAdr"`
		RprsTxprDclsNm         string `xml:"rprsTxprDclsNm"`
		AdrId                  string `xml:"adrId"`
		CrpClCd                string `xml:"crpClCd"`
		RprsRoadSggNm          string `xml:"rprsRoadSggNm"`
		WhtxDutyClNm           string `xml:"whtxDutyClNm"`
		HofcYn                 string `xml:"hofcYn"`
		Cpamt                  string `xml:"cpamt"`
		UndrBldClCd            string `xml:"undrBldClCd"`
		RoadSidoNm             string `xml:"roadSidoNm"`
		CrpKndCd               string `xml:"crpKndCd"`
		AsceYn                 string `xml:"asceYn"`
		RprsEnglRoadNm         string `xml:"rprsEnglRoadNm"`
		HofcRprsTxprDscmDt     string `xml:"hofcRprsTxprDscmDt"`
		SggCd                  string `xml:"sggCd"`
		RprsRoadYmdgNm         string `xml:"rprsRoadYmdgNm"`
		RprsBldSnoAdr          string `xml:"rprsBldSnoAdr"`
		RprsLdAdr              string `xml:"rprsLdAdr"`
		TxprStatNm             string `xml:"txprStatNm"`
		EdctxYn                string `xml:"edctxYn"`
		RoadBscAdr             string `xml:"roadBscAdr"`
		ProfBsBgnnDt           string `xml:"profBsBgnnDt"`
		BldBlckAdr             string `xml:"bldBlckAdr"`
		ScntxTxtnClNm          string `xml:"scntxTxtnClNm"`
		RprsHoAdr              string `xml:"rprsHoAdr"`
		RprsAdrId              string `xml:"rprsAdrId"`
		SidoEnglNm             string `xml:"sidoEnglNm"`
		EcstParCft             string `xml:"ecstParCft"`
		RoadSggNm              string `xml:"roadSggNm"`
		HofcRprsTin            string `xml:"hofcRprsTin"`
		MemNo                  string `xml:"memNo"`
		TxprClsfNm             string `xml:"txprClsfNm"`
		BmanClCd               string `xml:"bmanClCd"`
		FrsRgtDtm              string `xml:"frsRgtDtm"`
		AsctCrpClCd            string `xml:"asctCrpClCd"`
		LdZip                  string `xml:"ldZip"`
		WhlPmtYn               string `xml:"whlPmtYn"`
		RprsYmdgNm             string `xml:"rprsYmdgNm"`
		RprsSggCd              string `xml:"rprsSggCd"`
		PfbAnhsSfl             string `xml:"pfbAnhsSfl"`
		HofcRprsSgleTxprNm     string `xml:"hofcRprsSgleTxprNm"`
		BldSnoAdr              string `xml:"bldSnoAdr"`
		BldPmnoAdr             string `xml:"bldPmnoAdr"`
		RprsRoadSidoNm         string `xml:"rprsRoadSidoNm"`
		RpnAdr                 string `xml:"rpnAdr"`
		StatusValue            string `xml:"statusValue"`
		RprsResno              string `xml:"rprsResno"`
		Fnm                    string `xml:"fnm"`
		CntrAstsYn             string `xml:"cntrAstsYn"`
		TfbCd                  string `xml:"tfbCd"`
		LdBscAdr               string `xml:"ldBscAdr"`
		YmdgNm                 string `xml:"ymdgNm"`
		AsctCrpPubCrpYn        string `xml:"asctCrpPubCrpYn"`
		AbftChrgDutsNm         string `xml:"abftChrgDutsNm"`
		RprsRoadBscAdr         string `xml:"rprsRoadBscAdr"`
		JntBmanYn              string `xml:"jntBmanYn"`
		RprsBldBlckAdr         string `xml:"rprsBldBlckAdr"`
		BsyrEndMd              string `xml:"bsyrEndMd"`
		RprsLdZip              string `xml:"rprsLdZip"`
		RprsBldDnadr           string `xml:"rprsBldDnadr"`
		RprsTxprDclsCd         string `xml:"rprsTxprDclsCd"`
		RprsSidoEnglNm         string `xml:"rprsSidoEnglNm"`
		RoadSggZip             string `xml:"roadSggZip"`
		RprsLdBscAdr           string `xml:"rprsLdBscAdr"`
		WhtxDutyClCd           string `xml:"whtxDutyClCd"`
		CrpId                  string `xml:"crpId"`
		RprsRoadNm             string `xml:"rprsRoadNm"`
		RprsSidoNm             string `xml:"rprsSidoNm"`
		RcpnDt                 string `xml:"rcpnDt"`
		BunjAdr                string `xml:"bunjAdr"`
		RprsBunjAdr            string `xml:"rprsBunjAdr"`
		HofcNm                 string `xml:"hofcNm"`
		BmanRgtcPybClCd        string `xml:"bmanRgtcPybClCd"`
		HofcRprsTxprNm         string `xml:"hofcRprsTxprNm"`
		CrpDscmDt              string `xml:"crpDscmDt"`
		BldHoAdr               string `xml:"bldHoAdr"`
		RprsBldPmnoAdr         string `xml:"rprsBldPmnoAdr"`
		CrpChrcNm              string `xml:"crpChrcNm"`
		AbftChrgTxhfOgzNm      string `xml:"abftChrgTxhfOgzNm"`
		CmTxhfOgzCd            string `xml:"cmTxhfOgzCd"`
		BmanRgtDt              string `xml:"bmanRgtDt"`
		CrpTin                 string `xml:"crpTin"`
		HofcRprsResno          string `xml:"hofcRprsResno"`
		RoadNm                 string `xml:"roadNm"`
		SidoNm                 string `xml:"sidoNm"`
		RprsMp                 string `xml:"rprsMp"`
		ScntxTxtnClCd          string `xml:"scntxTxtnClCd"`
		YmdgZip                string `xml:"ymdgZip"`
		RprsYmdgZip            string `xml:"rprsYmdgZip"`
		RoadDadr               string `xml:"roadDadr"`
		RprsBldHoAdr           string `xml:"rprsBldHoAdr"`
		Count                  string `xml:"count"`
		PfbMhSfl               string `xml:"pfbMhSfl"`
		CrpStatCd              string `xml:"crpStatCd"`
		RprsHmTelno            string `xml:"rprsHmTelno"`
		RprsTin                string `xml:"rprsTin"`
		RprsSpcaCd             string `xml:"rprsSpcaCd"`
		Zip                    string `xml:"zip"`
		BsyrStrtMd             string `xml:"bsyrStrtMd"`
		YmdgEnglNm             string `xml:"ymdgEnglNm"`
		RprsRoadDadr           string `xml:"rprsRoadDadr"`
		RprsRoadSggZip         string `xml:"rprsRoadSggZip"`
		CrpTxprDclsCd          string `xml:"crpTxprDclsCd"`
		RprsTxprDscmDt         string `xml:"rprsTxprDscmDt"`
		RprsYmdgSn             string `xml:"rprsYmdgSn"`
		NcoTypeCd              string `xml:"ncoTypeCd"`
		RprsYmdgCd             string `xml:"rprsYmdgCd"`
		SpcaCd                 string `xml:"spcaCd"`
		RprsLdDadr             string `xml:"rprsLdDadr"`
		ProfBsTfbCd            string `xml:"profBsTfbCd"`
		RprsRoadYmdgZip        string `xml:"rprsRoadYmdgZip"`
		PfbTelno               string `xml:"pfbTelno"`
		YmdgSn                 string `xml:"ymdgSn"`
		YmdgCd                 string `xml:"ymdgCd"`
		RprsTxprStatCd         string `xml:"rprsTxprStatCd"`
		PfbPsenClNm            string `xml:"pfbPsenClNm"`
		TotaPblStocCnt         string `xml:"totaPblStocCnt"`
		Crpno                  string `xml:"crpno"`
		LdDadr                 string `xml:"ldDadr"`
		AprsCpamt              string `xml:"aprsCpamt"`
		PjpYn                  string `xml:"pjpYn"`
		RprsSggZip             string `xml:"rprsSggZip"`
		RoadAdr                string `xml:"roadAdr"`
		RprsRoadAdr            string `xml:"rprsRoadAdr"`
		RprsSgleTxprNm         string `xml:"rprsSgleTxprNm"`
		PjpCmntDnoCd           string `xml:"pjpCmntDnoCd"`
		CrpNm                  string `xml:"crpNm"`
		RprsSidoCd             string `xml:"rprsSidoCd"`
		RprsTxprClsfCd         string `xml:"rprsTxprClsfCd"`
		RprsYmdgEnglNm         string `xml:"rprsYmdgEnglNm"`
		CrpClNm                string `xml:"crpClNm"`
		InvmAmt                string `xml:"invmAmt"`
		SggZip                 string `xml:"sggZip"`
		MmrAmt                 string `xml:"mmrAmt"`
		RprsLdCd               string `xml:"rprsLdCd"`
		CrpChrcCd              string `xml:"crpChrcCd"`
		CrpKndNm               string `xml:"crpKndNm"`
		LstAltDtm              string `xml:"lstAltDtm"`
		RprsZip                string `xml:"rprsZip"`
		SidoCd                 string `xml:"sidoCd"`
		RprsTxprDscmNoClCd     string `xml:"rprsTxprDscmNoClCd"`
		PfbSfl                 string `xml:"pfbSfl"`
		SggNm                  string `xml:"sggNm"`
		RprsUndrBldClCd        string `xml:"rprsUndrBldClCd"`
		RoadNmCd               string `xml:"roadNmCd"`
		LstYn                  string `xml:"lstYn"`
		RprsTxprNm             string `xml:"rprsTxprNm"`
		RoadYmdgZip            string `xml:"roadYmdgZip"`
		TxprDclsNm             string `xml:"txprDclsNm"`
		RprsTxprDscmNoDplcClCd string `xml:"rprsTxprDscmNoDplcClCd"`
		GpBsYn                 string `xml:"gpBsYn"`
		CrdcJnnYn              string `xml:"crdcJnnYn"`
		RprsMaceCd             string `xml:"rprsMaceCd"`
		PuinKndCd              string `xml:"puinKndCd"`
		RprsRoadNmCd           string `xml:"rprsRoadNmCd"`
		BcNm                   string `xml:"bcNm"`
		BmanClNm               string `xml:"bmanClNm"`
		RprsSggEnglNm          string `xml:"rprsSggEnglNm"`
		RoadZip                string `xml:"roadZip"`
		YmdgClCd               string `xml:"ymdgClCd"`
		RprsRoadZip            string `xml:"rprsRoadZip"`
		CnctTxpYn              string `xml:"cnctTxpYn"`
		PutaYn                 string `xml:"putaYn"`
		TxpNoAprvTrgtCrpYn     string `xml:"txpNoAprvTrgtCrpYn"`
		EnglRoadNm             string `xml:"englRoadNm"`
		CrpCmpClCd             string `xml:"crpCmpClCd"`
		HofcRprsTxprDclsCd     string `xml:"hofcRprsTxprDclsCd"`
		RprsSggNm              string `xml:"rprsSggNm"`
		HofcTin                string `xml:"hofcTin"`
		HofcBsno               string `xml:"hofcBsno"`
		RoadYmdgNm             string `xml:"roadYmdgNm"`
		MaceCd                 string `xml:"maceCd"`
		PfbEml                 string `xml:"pfbEml"`
		LdCd                   string `xml:"ldCd"`
	} `xml:"map"`
	List []struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr"`
		Map  struct {
			Text                   string `xml:",chardata"`
			ID                     string `xml:"id,attr"`
			RprsYmdgClCd           string `xml:"rprsYmdgClCd"`
			ItmNm                  string `xml:"itmNm"`
			PfbPsenClCd            string `xml:"pfbPsenClCd"`
			PgBsTypeCd             string `xml:"pgBsTypeCd"`
			VatTxtnTrgtYn          string `xml:"vatTxtnTrgtYn"`
			Lfamt                  string `xml:"lfamt"`
			LdAdr                  string `xml:"ldAdr"`
			HofcTxprStatCd         string `xml:"hofcTxprStatCd"`
			SggEnglNm              string `xml:"sggEnglNm"`
			HofcRprsResnoDec       string `xml:"hofcRprsResnoDec"`
			EmplCnt                string `xml:"emplCnt"`
			OnsCpamt               string `xml:"onsCpamt"`
			HoAdr                  string `xml:"hoAdr"`
			RprsTxprDclsNm         string `xml:"rprsTxprDclsNm"`
			AdrId                  string `xml:"adrId"`
			CrpClCd                string `xml:"crpClCd"`
			RprsRoadSggNm          string `xml:"rprsRoadSggNm"`
			WhtxDutyClNm           string `xml:"whtxDutyClNm"`
			HofcYn                 string `xml:"hofcYn"`
			Cpamt                  string `xml:"cpamt"`
			UndrBldClCd            string `xml:"undrBldClCd"`
			RoadSidoNm             string `xml:"roadSidoNm"`
			CrpKndCd               string `xml:"crpKndCd"`
			AsceYn                 string `xml:"asceYn"`
			RprsEnglRoadNm         string `xml:"rprsEnglRoadNm"`
			HofcRprsTxprDscmDt     string `xml:"hofcRprsTxprDscmDt"`
			SggCd                  string `xml:"sggCd"`
			RprsRoadYmdgNm         string `xml:"rprsRoadYmdgNm"`
			RprsBldSnoAdr          string `xml:"rprsBldSnoAdr"`
			RprsLdAdr              string `xml:"rprsLdAdr"`
			TxprStatNm             string `xml:"txprStatNm"`
			EdctxYn                string `xml:"edctxYn"`
			RoadBscAdr             string `xml:"roadBscAdr"`
			ProfBsBgnnDt           string `xml:"profBsBgnnDt"`
			BldBlckAdr             string `xml:"bldBlckAdr"`
			ScntxTxtnClNm          string `xml:"scntxTxtnClNm"`
			RprsHoAdr              string `xml:"rprsHoAdr"`
			RprsAdrId              string `xml:"rprsAdrId"`
			TxprDclsCd             string `xml:"txprDclsCd"`
			TxprDscmNoDplcClCd     string `xml:"txprDscmNoDplcClCd"`
			SidoEnglNm             string `xml:"sidoEnglNm"`
			EcstParCft             string `xml:"ecstParCft"`
			RoadSggNm              string `xml:"roadSggNm"`
			HofcRprsTin            string `xml:"hofcRprsTin"`
			MemNo                  string `xml:"memNo"`
			TxprClsfNm             string `xml:"txprClsfNm"`
			BmanClCd               string `xml:"bmanClCd"`
			FrsRgtDtm              string `xml:"frsRgtDtm"`
			AsctCrpClCd            string `xml:"asctCrpClCd"`
			LdZip                  string `xml:"ldZip"`
			WhlPmtYn               string `xml:"whlPmtYn"`
			RprsYmdgNm             string `xml:"rprsYmdgNm"`
			RprsSggCd              string `xml:"rprsSggCd"`
			PfbAnhsSfl             string `xml:"pfbAnhsSfl"`
			HofcRprsSgleTxprNm     string `xml:"hofcRprsSgleTxprNm"`
			BldSnoAdr              string `xml:"bldSnoAdr"`
			BldPmnoAdr             string `xml:"bldPmnoAdr"`
			Tin                    string `xml:"tin"`
			RprsRoadSidoNm         string `xml:"rprsRoadSidoNm"`
			RpnAdr                 string `xml:"rpnAdr"`
			StatusValue            string `xml:"statusValue"`
			RprsResno              string `xml:"rprsResno"`
			Fnm                    string `xml:"fnm"`
			CntrAstsYn             string `xml:"cntrAstsYn"`
			TfbCd                  string `xml:"tfbCd"`
			LdBscAdr               string `xml:"ldBscAdr"`
			YmdgNm                 string `xml:"ymdgNm"`
			AsctCrpPubCrpYn        string `xml:"asctCrpPubCrpYn"`
			AbftChrgDutsNm         string `xml:"abftChrgDutsNm"`
			RprsRoadBscAdr         string `xml:"rprsRoadBscAdr"`
			JntBmanYn              string `xml:"jntBmanYn"`
			RprsBldBlckAdr         string `xml:"rprsBldBlckAdr"`
			BsyrEndMd              string `xml:"bsyrEndMd"`
			RprsLdZip              string `xml:"rprsLdZip"`
			RprsBldDnadr           string `xml:"rprsBldDnadr"`
			TxprDscmNoClCd         string `xml:"txprDscmNoClCd"`
			RprsTxprDclsCd         string `xml:"rprsTxprDclsCd"`
			RprsSidoEnglNm         string `xml:"rprsSidoEnglNm"`
			RoadSggZip             string `xml:"roadSggZip"`
			TxprDscmDt             string `xml:"txprDscmDt"`
			RprsLdBscAdr           string `xml:"rprsLdBscAdr"`
			WhtxDutyClCd           string `xml:"whtxDutyClCd"`
			CrpId                  string `xml:"crpId"`
			RprsRoadNm             string `xml:"rprsRoadNm"`
			RprsSidoNm             string `xml:"rprsSidoNm"`
			RcpnDt                 string `xml:"rcpnDt"`
			BunjAdr                string `xml:"bunjAdr"`
			RprsBunjAdr            string `xml:"rprsBunjAdr"`
			HofcNm                 string `xml:"hofcNm"`
			BmanRgtcPybClCd        string `xml:"bmanRgtcPybClCd"`
			HofcRprsTxprNm         string `xml:"hofcRprsTxprNm"`
			CrpDscmDt              string `xml:"crpDscmDt"`
			BldHoAdr               string `xml:"bldHoAdr"`
			RprsBldPmnoAdr         string `xml:"rprsBldPmnoAdr"`
			CrpChrcNm              string `xml:"crpChrcNm"`
			AbftChrgTxhfOgzNm      string `xml:"abftChrgTxhfOgzNm"`
			CmTxhfOgzCd            string `xml:"cmTxhfOgzCd"`
			BmanRgtDt              string `xml:"bmanRgtDt"`
			CrpTin                 string `xml:"crpTin"`
			TxprStatCd             string `xml:"txprStatCd"`
			HofcRprsResno          string `xml:"hofcRprsResno"`
			RoadNm                 string `xml:"roadNm"`
			SidoNm                 string `xml:"sidoNm"`
			RprsMp                 string `xml:"rprsMp"`
			ScntxTxtnClCd          string `xml:"scntxTxtnClCd"`
			YmdgZip                string `xml:"ymdgZip"`
			RprsYmdgZip            string `xml:"rprsYmdgZip"`
			RoadDadr               string `xml:"roadDadr"`
			RprsBldHoAdr           string `xml:"rprsBldHoAdr"`
			Count                  string `xml:"count"`
			PfbMhSfl               string `xml:"pfbMhSfl"`
			CrpStatCd              string `xml:"crpStatCd"`
			RprsHmTelno            string `xml:"rprsHmTelno"`
			RprsTin                string `xml:"rprsTin"`
			RprsSpcaCd             string `xml:"rprsSpcaCd"`
			Zip                    string `xml:"zip"`
			TxprClsfCd             string `xml:"txprClsfCd"`
			BsyrStrtMd             string `xml:"bsyrStrtMd"`
			YmdgEnglNm             string `xml:"ymdgEnglNm"`
			RprsRoadDadr           string `xml:"rprsRoadDadr"`
			RprsRoadSggZip         string `xml:"rprsRoadSggZip"`
			CrpTxprDclsCd          string `xml:"crpTxprDclsCd"`
			RprsTxprDscmDt         string `xml:"rprsTxprDscmDt"`
			TxprMntgYn             string `xml:"txprMntgYn"`
			RprsYmdgSn             string `xml:"rprsYmdgSn"`
			NcoTypeCd              string `xml:"ncoTypeCd"`
			RprsYmdgCd             string `xml:"rprsYmdgCd"`
			SpcaCd                 string `xml:"spcaCd"`
			RprsLdDadr             string `xml:"rprsLdDadr"`
			ProfBsTfbCd            string `xml:"profBsTfbCd"`
			RprsRoadYmdgZip        string `xml:"rprsRoadYmdgZip"`
			PfbTelno               string `xml:"pfbTelno"`
			YmdgSn                 string `xml:"ymdgSn"`
			YmdgCd                 string `xml:"ymdgCd"`
			RprsTxprStatCd         string `xml:"rprsTxprStatCd"`
			AbftChrgDutsCd         string `xml:"abftChrgDutsCd"`
			PfbPsenClNm            string `xml:"pfbPsenClNm"`
			TotaPblStocCnt         string `xml:"totaPblStocCnt"`
			Crpno                  string `xml:"crpno"`
			LdDadr                 string `xml:"ldDadr"`
			AprsCpamt              string `xml:"aprsCpamt"`
			PjpYn                  string `xml:"pjpYn"`
			RprsSggZip             string `xml:"rprsSggZip"`
			RoadAdr                string `xml:"roadAdr"`
			RprsRoadAdr            string `xml:"rprsRoadAdr"`
			RprsSgleTxprNm         string `xml:"rprsSgleTxprNm"`
			PjpCmntDnoCd           string `xml:"pjpCmntDnoCd"`
			CrpNm                  string `xml:"crpNm"`
			RprsSidoCd             string `xml:"rprsSidoCd"`
			RprsTxprClsfCd         string `xml:"rprsTxprClsfCd"`
			RprsYmdgEnglNm         string `xml:"rprsYmdgEnglNm"`
			CrpClNm                string `xml:"crpClNm"`
			InvmAmt                string `xml:"invmAmt"`
			SggZip                 string `xml:"sggZip"`
			MmrAmt                 string `xml:"mmrAmt"`
			RprsLdCd               string `xml:"rprsLdCd"`
			AbftChrgTxhfOgzCd      string `xml:"abftChrgTxhfOgzCd"`
			CrpChrcCd              string `xml:"crpChrcCd"`
			CrpKndNm               string `xml:"crpKndNm"`
			LstAltDtm              string `xml:"lstAltDtm"`
			RprsZip                string `xml:"rprsZip"`
			SidoCd                 string `xml:"sidoCd"`
			RprsTxprDscmNoClCd     string `xml:"rprsTxprDscmNoClCd"`
			PfbSfl                 string `xml:"pfbSfl"`
			SggNm                  string `xml:"sggNm"`
			RprsUndrBldClCd        string `xml:"rprsUndrBldClCd"`
			RoadNmCd               string `xml:"roadNmCd"`
			LstYn                  string `xml:"lstYn"`
			RprsTxprNm             string `xml:"rprsTxprNm"`
			RoadYmdgZip            string `xml:"roadYmdgZip"`
			TxprDclsNm             string `xml:"txprDclsNm"`
			RprsTxprDscmNoDplcClCd string `xml:"rprsTxprDscmNoDplcClCd"`
			GpBsYn                 string `xml:"gpBsYn"`
			CrdcJnnYn              string `xml:"crdcJnnYn"`
			RprsMaceCd             string `xml:"rprsMaceCd"`
			PuinKndCd              string `xml:"puinKndCd"`
			RprsRoadNmCd           string `xml:"rprsRoadNmCd"`
			BcNm                   string `xml:"bcNm"`
			BmanClNm               string `xml:"bmanClNm"`
			RprsSggEnglNm          string `xml:"rprsSggEnglNm"`
			RoadZip                string `xml:"roadZip"`
			YmdgClCd               string `xml:"ymdgClCd"`
			RprsRoadZip            string `xml:"rprsRoadZip"`
			CnctTxpYn              string `xml:"cnctTxpYn"`
			TxprNm                 string `xml:"txprNm"`
			PutaYn                 string `xml:"putaYn"`
			TxpNoAprvTrgtCrpYn     string `xml:"txpNoAprvTrgtCrpYn"`
			EnglRoadNm             string `xml:"englRoadNm"`
			CrpCmpClCd             string `xml:"crpCmpClCd"`
			HofcRprsTxprDclsCd     string `xml:"hofcRprsTxprDclsCd"`
			RprsSggNm              string `xml:"rprsSggNm"`
			HofcTin                string `xml:"hofcTin"`
			HofcBsno               string `xml:"hofcBsno"`
			RoadYmdgNm             string `xml:"roadYmdgNm"`
			MaceCd                 string `xml:"maceCd"`
			PfbEml                 string `xml:"pfbEml"`
			LdCd                   string `xml:"ldCd"`
			TxprDscmNoEncCntn      string `xml:"txprDscmNoEncCntn"`
		} `xml:"map"`
	} `xml:"list"`
	RowCount string `xml:"rowCount"`
	Tin      string `xml:"tin"`
}

func (c *Client) CorpInfo(corpNum, tin, userClCd string) (corpInfoRes *CorpInfoResponse, err error) {
	res, err := req.Post(
		"https://teht.hometax.go.kr/wqAction.do?actionId=ATTABZAA001R17&screenId=UTEABGAA21&popupYn=false&realScreenId=",
		c.httpCli,
		req.BodyXML(fmt.Sprintf(`<map id="ATTABZAA001R17">
<tin>%s</tin>
<txprClsfCd>02</txprClsfCd>
<txprDscmNo>%s</txprDscmNo>
<txprDscmNoClCd/>
<txprDscmDt/>
<searchOrder>02/01</searchOrder>
<outDes>bmanBscInfrInqrDVO</outDes>
<txprNm/>
<crpTin/>
<mntgTxprIcldYn/>
<resnoAltHstrInqrYn/>
<resnoAltHstrInqrBaseDtm/>
</map><nts<nts>nts>`, tin, corpNum)+aefire.RandString(44)))

	if aefire.LogIfError(err) {
		return nil, errors.New("사용자 정보 조회가 실패했습니다 : 홈택스 연결 오류 발생")
	}

	corpInfoRes = &CorpInfoResponse{}
	if aefire.LogIfError(res.ToXML(corpInfoRes)) {
		return nil, errors.New("사용자 정보 조회가 실패했습니다 : 홈택스 응답 해석 실패")
	}

	if corpInfoRes.Map[0].Result != "S" {
		if corpInfoRes.Map[0].Msg != "" {
			return nil, errors.New("사용자 정보 조회가 실패했습니다 : " + corpInfoRes.Map[0].Msg)
		} else {
			//println(res.String())
			return nil, errors.New("사용자 정보 조회가 실패했습니다 : 홈택스 요청 오류 발생")
		}
	}

	return
}
