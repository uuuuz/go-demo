import hmac

signature_payload = '1644458608000GET/wallet/balances'
signature_payload = signature_payload.encode()
signature = hmac.new('AL3Tp_vA1On6f188t0MsDoYr9RHqpmfkLN7H3w2q'.encode(), signature_payload, 'sha256').hexdigest()

print(signature)