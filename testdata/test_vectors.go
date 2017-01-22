package testdata

const (
	//Expected unmarshaled values
	TEST_REALM                  = "ATHENA.MIT.EDU"
	TEST_CIPHERTEXT             = "krbASN.1 test message"
	TEST_TIME_FORMAT            = "20060102150405"
	TEST_TIME                   = "19940610060317"
	TEST_PRINCIPALNAME_NAMETYPE = 1
	TEST_KVNO                   = 5
	TEST_ETYPE                  = 0
	TEST_NONCE                  = 42
	TEST_AUTHORIZATION_DATA_TYPE = 1
	TEST_AUTHORIZATION_DATA_VALUE = "foobar"
	TEST_PADATA_TYPE = 13
	TEST_PADATA_VALUE = "pa-data"
)

var TEST_PRINCIPALNAME_NAMESTRING = []string{"hftsai", "extra"}

//The test vectors have been sourced from https://github.com/krb5/krb5/blob/master/src/tests/asn.1/reference_encode.out
var TestVectors = map[string]string{
	"encode_krb5_authenticator":                                  "6281A130819EA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A30F300DA003020101A106040431323334A405020301E240A511180F31393934303631303036303331375AA6133011A003020101A10A04083132333435363738A703020111A8243022300FA003020101A1080406666F6F626172300FA003020101A1080406666F6F626172",
	"encode_krb5_authenticator(optionalsempty)":                  "624F304DA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A405020301E240A511180F31393934303631303036303331375A",
	"encode_krb5_authenticator(optionalsNULL)":                   "624F304DA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A405020301E240A511180F31393934303631303036303331375A",
	"encode_krb5_ticket":                                         "615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	"encode_krb5_keyblock":                                       "3011A003020101A10A04083132333435363738",
	"encode_krb5_enc_tkt_part":                                   "6382011430820110A007030500FEDCBA98A1133011A003020101A10A04083132333435363738A2101B0E415448454E412E4D49542E454455A31A3018A003020101A111300F1B066866747361691B056578747261A42E302CA003020101A12504234544552C4D49542E2C415448454E412E2C57415348494E47544F4E2E4544552C43532EA511180F31393934303631303036303331375AA611180F31393934303631303036303331375AA711180F31393934303631303036303331375AA811180F31393934303631303036303331375AA920301E300DA003020102A106040412D00023300DA003020102A106040412D00023AA243022300FA003020101A1080406666F6F626172300FA003020101A1080406666F6F626172",
	"encode_krb5_enc_tkt_part(optionalsNULL)":                    "6381A53081A2A007030500FEDCBA98A1133011A003020101A10A04083132333435363738A2101B0E415448454E412E4D49542E454455A31A3018A003020101A111300F1B066866747361691B056578747261A42E302CA003020101A12504234544552C4D49542E2C415448454E412E2C57415348494E47544F4E2E4544552C43532EA511180F31393934303631303036303331375AA711180F31393934303631303036303331375A",
	"encode_krb5_enc_kdc_rep_part":                               "7A82010E3082010AA0133011A003020101A10A04083132333435363738A13630343018A0030201FBA111180F31393934303631303036303331375A3018A0030201FBA111180F31393934303631303036303331375AA20302012AA311180F31393934303631303036303331375AA407030500FEDCBA98A511180F31393934303631303036303331375AA611180F31393934303631303036303331375AA711180F31393934303631303036303331375AA811180F31393934303631303036303331375AA9101B0E415448454E412E4D49542E454455AA1A3018A003020101A111300F1B066866747361691B056578747261AB20301E300DA003020102A106040412D00023300DA003020102A106040412D00023",
	"encode_krb5_enc_kdc_rep_part(optionalsNULL)":                "7A81B23081AFA0133011A003020101A10A04083132333435363738A13630343018A0030201FBA111180F31393934303631303036303331375A3018A0030201FBA111180F31393934303631303036303331375AA20302012AA407030500FE5CBA98A511180F31393934303631303036303331375AA711180F31393934303631303036303331375AA9101B0E415448454E412E4D49542E454455AA1A3018A003020101A111300F1B066866747361691B056578747261",
	"encode_krb5_as_rep":                                         "6B81EA3081E7A003020105A10302010BA22630243010A10302010DA209040770612D646174613010A10302010DA209040770612D64617461A3101B0E415448454E412E4D49542E454455A41A3018A003020101A111300F1B066866747361691B056578747261A55E615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765A6253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	"encode_krb5_as_rep(optionalsNULL)":                          "6B81C23081BFA003020105A10302010BA3101B0E415448454E412E4D49542E454455A41A3018A003020101A111300F1B066866747361691B056578747261A55E615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765A6253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	"encode_krb5_tgs_rep":                                        "6D81EA3081E7A003020105A10302010DA22630243010A10302010DA209040770612D646174613010A10302010DA209040770612D64617461A3101B0E415448454E412E4D49542E454455A41A3018A003020101A111300F1B066866747361691B056578747261A55E615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765A6253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	"encode_krb5_tgs_rep(optionalsNULL)":                         "6D81C23081BFA003020105A10302010DA3101B0E415448454E412E4D49542E454455A41A3018A003020101A111300F1B066866747361691B056578747261A55E615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765A6253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	"encode_krb5_ap_req":                                         "6E819D30819AA003020105A10302010EA207030500FEDCBA98A35E615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765A4253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	//"encode_krb5_ap_rep":                                         "6F333031A003020105A10302010FA2253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	//"encode_krb5_ap_rep_enc_part":                                "7B363034A011180F31393934303631303036303331375AA105020301E240A2133011A003020101A10A04083132333435363738A303020111",
	//"encode_krb5_ap_rep_enc_part(optionalsNULL)":                 "7B1C301AA011180F31393934303631303036303331375AA105020301E240",
	"encode_krb5_as_req":                                         "6A8201E4308201E0A103020105A20302010AA32630243010A10302010DA209040770612D646174613010A10302010DA209040770612D64617461A48201AA308201A6A007030500FEDCBA90A11A3018A003020101A111300F1B066866747361691B056578747261A2101B0E415448454E412E4D49542E454455A31A3018A003020101A111300F1B066866747361691B056578747261A411180F31393934303631303036303331375AA511180F31393934303631303036303331375AA611180F31393934303631303036303331375AA70302012AA8083006020100020101A920301E300DA003020102A106040412D00023300DA003020102A106040412D00023AA253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765AB81BF3081BC615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	"encode_krb5_as_req(optionalsNULLexceptsecond_ticket)":       "6A82011430820110A103020105A20302010AA48201023081FFA007030500FEDCBA98A2101B0E415448454E412E4D49542E454455A511180F31393934303631303036303331375AA70302012AA8083006020100020101AB81BF3081BC615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	"encode_krb5_as_req(optionalsNULLexceptserver)":              "6A693067A103020105A20302010AA45B3059A007030500FEDCBA90A2101B0E415448454E412E4D49542E454455A31A3018A003020101A111300F1B066866747361691B056578747261A511180F31393934303631303036303331375AA70302012AA8083006020100020101",
	"encode_krb5_tgs_req":                                        "6C8201E4308201E0A103020105A20302010CA32630243010A10302010DA209040770612D646174613010A10302010DA209040770612D64617461A48201AA308201A6A007030500FEDCBA90A11A3018A003020101A111300F1B066866747361691B056578747261A2101B0E415448454E412E4D49542E454455A31A3018A003020101A111300F1B066866747361691B056578747261A411180F31393934303631303036303331375AA511180F31393934303631303036303331375AA611180F31393934303631303036303331375AA70302012AA8083006020100020101A920301E300DA003020102A106040412D00023300DA003020102A106040412D00023AA253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765AB81BF3081BC615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	"encode_krb5_tgs_req(optionalsNULLexceptsecond_ticket)":      "6C82011430820110A103020105A20302010CA48201023081FFA007030500FEDCBA98A2101B0E415448454E412E4D49542E454455A511180F31393934303631303036303331375AA70302012AA8083006020100020101AB81BF3081BC615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	"encode_krb5_tgs_req(optionalsNULLexceptserver)":             "6C693067A103020105A20302010CA45B3059A007030500FEDCBA90A2101B0E415448454E412E4D49542E454455A31A3018A003020101A111300F1B066866747361691B056578747261A511180F31393934303631303036303331375AA70302012AA8083006020100020101",
	"encode_krb5_kdc_req_body":                                   "308201A6A007030500FEDCBA90A11A3018A003020101A111300F1B066866747361691B056578747261A2101B0E415448454E412E4D49542E454455A31A3018A003020101A111300F1B066866747361691B056578747261A411180F31393934303631303036303331375AA511180F31393934303631303036303331375AA611180F31393934303631303036303331375AA70302012AA8083006020100020101A920301E300DA003020102A106040412D00023300DA003020102A106040412D00023AA253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765AB81BF3081BC615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	"encode_krb5_kdc_req_body(optionalsNULLexceptsecond_ticket)": "3081FFA007030500FEDCBA98A2101B0E415448454E412E4D49542E454455A511180F31393934303631303036303331375AA70302012AA8083006020100020101AB81BF3081BC615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	"encode_krb5_kdc_req_body(optionalsNULLexceptserver)":        "3059A007030500FEDCBA90A2101B0E415448454E412E4D49542E454455A31A3018A003020101A111300F1B066866747361691B056578747261A511180F31393934303631303036303331375AA70302012AA8083006020100020101",
	//"encode_krb5_safe":                                           "746E306CA003020105A103020114A24F304DA00A04086B72623564617461A111180F31393934303631303036303331375AA205020301E240A303020111A40F300DA003020102A106040412D00023A50F300DA003020102A106040412D00023A30F300DA003020101A106040431323334",
	//"encode_krb5_safe(optionalsNULL)":                            "743E303CA003020105A103020114A21F301DA00A04086B72623564617461A40F300DA003020102A106040412D00023A30F300DA003020101A106040431323334",
	//"encode_krb5_priv":                                           "75333031A003020105A103020115A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	//"encode_krb5_enc_priv_part":                                  "7C4F304DA00A04086B72623564617461A111180F31393934303631303036303331375AA205020301E240A303020111A40F300DA003020102A106040412D00023A50F300DA003020102A106040412D00023",
	//"encode_krb5_enc_priv_part(optionalsNULL)":                   "7C1F301DA00A04086B72623564617461A40F300DA003020102A106040412D00023",
	//"encode_krb5_cred":                                           "7681F63081F3A003020105A103020116A281BF3081BC615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	//"encode_krb5_enc_cred_part":                                  "7D8202233082021FA08201DA308201D63081E8A0133011A003020101A10A04083132333435363738A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A307030500FEDCBA98A411180F31393934303631303036303331375AA511180F31393934303631303036303331375AA611180F31393934303631303036303331375AA711180F31393934303631303036303331375AA8101B0E415448454E412E4D49542E454455A91A3018A003020101A111300F1B066866747361691B056578747261AA20301E300DA003020102A106040412D00023300DA003020102A106040412D000233081E8A0133011A003020101A10A04083132333435363738A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A307030500FEDCBA98A411180F31393934303631303036303331375AA511180F31393934303631303036303331375AA611180F31393934303631303036303331375AA711180F31393934303631303036303331375AA8101B0E415448454E412E4D49542E454455A91A3018A003020101A111300F1B066866747361691B056578747261AA20301E300DA003020102A106040412D00023300DA003020102A106040412D00023A10302012AA211180F31393934303631303036303331375AA305020301E240A40F300DA003020102A106040412D00023A50F300DA003020102A106040412D00023",
	//"encode_krb5_enc_cred_part(optionalsNULL)":                   "7D82010E3082010AA0820106308201023015A0133011A003020101A10A040831323334353637383081E8A0133011A003020101A10A04083132333435363738A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A307030500FEDCBA98A411180F31393934303631303036303331375AA511180F31393934303631303036303331375AA611180F31393934303631303036303331375AA711180F31393934303631303036303331375AA8101B0E415448454E412E4D49542E454455A91A3018A003020101A111300F1B066866747361691B056578747261AA20301E300DA003020102A106040412D00023300DA003020102A106040412D00023",
	//"encode_krb5_error":                                          "7E81BA3081B7A003020105A10302011EA211180F31393934303631303036303331375AA305020301E240A411180F31393934303631303036303331375AA505020301E240A60302013CA7101B0E415448454E412E4D49542E454455A81A3018A003020101A111300F1B066866747361691B056578747261A9101B0E415448454E412E4D49542E454455AA1A3018A003020101A111300F1B066866747361691B056578747261AB0A1B086B72623564617461AC0A04086B72623564617461",
	//"encode_krb5_error(optionalsNULL)":                           "7E60305EA003020105A10302011EA305020301E240A411180F31393934303631303036303331375AA505020301E240A60302013CA9101B0E415448454E412E4D49542E454455AA1A3018A003020101A111300F1B066866747361691B056578747261",
	"encode_krb5_authorization_data":                             "3022300FA003020101A1080406666F6F626172300FA003020101A1080406666F6F626172",
	"encode_krb5_padata_sequence":                                "30243010A10302010DA209040770612D646174613010A10302010DA209040770612D64617461",
	//"encode_krb5_typed_data":                                     "30243010A00302010DA109040770612D646174613010A00302010DA109040770612D64617461",
	"encode_krb5_padata_sequence(empty)":                         "3000",
	"encode_krb5_etype_info":                                     "30333014A003020100A10D040B4D6F72746F6E27732023303005A0030201013014A003020102A10D040B4D6F72746F6E2773202332",
	"encode_krb5_etype_info(only1)":                              "30163014A003020100A10D040B4D6F72746F6E2773202330",
	"encode_krb5_etype_info(noinfo)":                             "3000",
	"encode_krb5_etype_info2":                                    "3051301EA003020100A10D1B0B4D6F72746F6E2773202330A208040673326B3A2030300FA003020101A208040673326B3A2031301EA003020102A10D1B0B4D6F72746F6E2773202332A208040673326B3A2032",
	"encode_krb5_etype_info2(only1)":                             "3020301EA003020100A10D1B0B4D6F72746F6E2773202330A208040673326B3A2030",
	"encode_krb5_pa_enc_ts":                                      "301AA011180F31393934303631303036303331375AA105020301E240",
	"encode_krb5_pa_enc_ts(nousec)":                              "3013A011180F31393934303631303036303331375A",
	"encode_krb5_enc_data":                                       "3023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	"encode_krb5_enc_data(MSB-setkvno)":                          "3026A003020100A1060204FF000000A21704156B726241534E2E312074657374206D657373616765",
	"encode_krb5_enc_data(kvno= -1)":                             "3023A003020100A1030201FFA21704156B726241534E2E312074657374206D657373616765",
	//"encode_krb5_sam_challenge_2":                                "3022A00D300B04096368616C6C656E6765A111300F300DA003020101A106040431323334",
	//"encode_krb5_sam_challenge_2_body":                           "3064A00302012AA10703050080000000A20B040974797065206E616D65A411040F6368616C6C656E6765206C6162656CA510040E6368616C6C656E67652069707365A6160414726573706F6E73655F70726F6D70742069707365A8050203543210A903020101",
	//"encode_krb5_sam_response_2":                                 "3042A00302012BA10703050080000000A20C040A747261636B2064617461A31D301BA003020101A10402020D36A20E040C6E6F6E6365206F7220736164A4050203543210",
	//"encode_krb5_enc_sam_response_enc_2":                         "301FA003020158A1180416656E635F73616D5F726573706F6E73655F656E635F32",
	//"encode_krb5_pa_for_user":                                    "304BA01A3018A003020101A111300F1B066866747361691B056578747261A1101B0E415448454E412E4D49542E454455A20F300DA003020101A106040431323334A30A1B086B72623564617461",
	//"encode_krb5_pa_s4u_x509_user":                               "3068A0553053A006020400CA149AA11A3018A003020101A111300F1B066866747361691B056578747261A2101B0E415448454E412E4D49542E454455A312041070615F7334755F783530395F75736572A40703050080000000A10F300DA003020101A106040431323334",
	"encode_krb5_ad_kdcissued":                                   "3065A00F300DA003020101A106040431323334A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3243022300FA003020101A1080406666F6F626172300FA003020101A1080406666F6F626172",
	//"encode_krb5_ad_signedpath_data":                             "3081C7A030302EA01A3018A003020101A111300F1B066866747361691B056578747261A1101B0E415448454E412E4D49542E454455A111180F31393934303631303036303331375AA2323030302EA01A3018A003020101A111300F1B066866747361691B056578747261A1101B0E415448454E412E4D49542E454455A32630243010A10302010DA209040770612D646174613010A10302010DA209040770612D64617461A4243022300FA003020101A1080406666F6F626172300FA003020101A1080406666F6F626172",
	//"encode_krb5_ad_signedpath":                                  "303EA003020101A10F300DA003020101A106040431323334A32630243010A10302010DA209040770612D646174613010A10302010DA209040770612D64617461",
	//"encode_krb5_iakerb_header":                                  "3018A10A04086B72623564617461A20A04086B72623564617461",
	//"encode_krb5_iakerb_finished":                                "3011A10F300DA003020101A106040431323334",
	//"encode_krb5_fast_response":                                  "30819FA02630243010A10302010DA209040770612D646174613010A10302010DA209040770612D64617461A1133011A003020101A10A04083132333435363738A25B3059A011180F31393934303631303036303331375AA105020301E240A2101B0E415448454E412E4D49542E454455A31A3018A003020101A111300F1B066866747361691B056578747261A40F300DA003020101A106040431323334A30302012A",
	//"encode_krb5_pa_fx_fast_reply":                               "A0293027A0253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	//"encode_krb5_otp_tokeninfo(optionalsNULL)":                   "300780050000000000",
	//"encode_krb5_otp_tokeninfo":                                  "307280050077000000810B4578616D706C65636F727082056861726B2183010A8401028509796F7572746F6B656E862875726E3A696574663A706172616D733A786D6C3A6E733A6B657970726F763A70736B633A686F7470A716300B0609608648016503040201300706052B0E03021A880203E8",
	//"encode_krb5_pa_otp_challenge(optionalsNULL)":                "301580086D696E6E6F6E6365A209300780050000000000",
	//"encode_krb5_pa_otp_challenge":                               "3081A580086D61786E6F6E6365810B7465737473657276696365A27D300780050000000000307280050077000000810B4578616D706C65636F727082056861726B2183010A8401028509796F7572746F6B656E862875726E3A696574663A706172616D733A786D6C3A6E733A6B657970726F763A70736B633A686F7470A716300B0609608648016503040201300706052B0E03021A880203E883076B657973616C74840431323334",
	//"encode_krb5_pa_otp_req(optionalsNULL)":                      "302C80050000000000A223A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	//"encode_krb5_pa_otp_req":                                     "3081B98005006000000081056E6F6E6365A223A003020100A103020105A21704156B726241534E2E312074657374206D657373616765A30B0609608648016503040201840203E8850566726F6773860A6D79666972737470696E87056861726B21880F31393934303631303036303331375A89033334368A01028B09796F7572746F6B656E8C2875726E3A696574663A706172616D733A786D6C3A6E733A6B657970726F763A70736B633A686F74708D0B4578616D706C65636F7270",
	//"encode_krb5_pa_otp_enc_req":                                 "300A80086B72623564617461",
	//"encode_krb5_kkdcp_message":                                  "308201FCA08201EC048201E86A8201E4308201E0A103020105A20302010AA32630243010A10302010DA209040770612D646174613010A10302010DA209040770612D64617461A48201AA308201A6A007030500FEDCBA98A11A3018A003020101A111300F1B066866747361691B056578747261A2101B0E415448454E412E4D49542E454455A31A3018A003020101A111300F1B066866747361691B056578747261A411180F31393934303631303036303331375AA511180F31393934303631303036303331375AA611180F31393934303631303036303331375AA70302012AA8083006020100020101A920301E300DA003020102A106040412D00023300DA003020102A106040412D00023AA253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765AB81BF3081BC615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765A10A1B086B72623564617461",
	//"encode_krb5_cammac(optionalsNULL)":                          "3012A010300E300CA003020101A1050403616431",
	//"encode_krb5_cammac":                                         "3081F2A01E301C300CA003020101A1050403616431300CA003020102A1050403616432A13D303BA01A3018A003020101A111300F1B066866747361691B056578747261A103020105A203020110A3133011A003020101A10A0408636B73756D6B6463A23D303BA01A3018A003020101A111300F1B066866747361691B056578747261A103020105A203020110A3133011A003020101A10A0408636B73756D737663A35230503013A311300FA003020101A1080406636B73756D313039A01A3018A003020101A111300F1B066866747361691B056578747261A103020105A203020110A311300FA003020101A1080406636B73756D32",
	//"encode_krb5_secure_cookie":                                  "302C02042DF8022530243010A10302010DA209040770612D646174613010A10302010DA209040770612D64617461",
}
