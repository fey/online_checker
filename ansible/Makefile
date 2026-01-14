ping:
	ansible all -i hosts -m ping -u root

ssh:
	ssh root@`cat hosts`

deploy:
	ansible-playbook deploy.yml -i hosts --vault-password-file=../.vault_password

fetch:
	ansible-playbook fetch.yml -i hosts --vault-password-file=../.vault_password

encrypt:
	ansible-vault encrypt --vault-password-file=../.vault_password group_vars/all/vault.yml

decrypt:
	ansible-vault decrypt --vault-password-file=../.vault_password group_vars/all/vault.yml
