(window.webpackJsonp=window.webpackJsonp||[]).push([[3],{356:function(e,t,r){e.exports=r.p+"assets/img/k8s-vault-webhook-logo.5976f5eb.svg"},357:function(e,t,r){e.exports=r.p+"assets/img/k8s-vault-webhook-arc-vault.5c511a3e.png"},358:function(e,t,r){e.exports=r.p+"assets/img/k8s-vault-webhook-arc-aws.5703a981.png"},375:function(e,t,r){"use strict";r.r(t);var a=r(45),n=Object(a.a)({},(function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[a("h1",{attrs:{id:"introduction"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#introduction"}},[e._v("#")]),e._v(" Introduction")]),e._v(" "),a("div",{attrs:{align:"center"}},[a("img",{attrs:{src:r(356),height:"120",width:"120"}})]),e._v(" "),a("p",[e._v("k8s-vault-webhook is a Kubernetes admission webhook which listen for the events related to Kubernetes resources for injecting secret directly from secret manager to pod, secret, and configmap.\nThe motive of creating this project is to provide a dynamic secret injection to containers/pods running inside Kubernetes from different secret managers for enhanced security.")]),e._v(" "),a("p",[e._v("Documentation is available here:- "),a("a",{attrs:{href:"https://ot-container-kit.github.io/k8s-vault-webhook/",target:"_blank",rel:"noopener noreferrer"}},[e._v("https://ot-container-kit.github.io/k8s-vault-webhook/"),a("OutboundLink")],1)]),e._v(" "),a("p",[e._v("The secret managers which are currently supported:-")]),e._v(" "),a("ul",[a("li",[a("strong",[a("a",{attrs:{href:"https://www.vaultproject.io/",target:"_blank",rel:"noopener noreferrer"}},[e._v("Hashicorp Vault"),a("OutboundLink")],1)])]),e._v(" "),a("li",[a("strong",[a("a",{attrs:{href:"https://aws.amazon.com/secrets-manager/",target:"_blank",rel:"noopener noreferrer"}},[e._v("AWS Secret Manager"),a("OutboundLink")],1)])])]),e._v(" "),a("p",[e._v("There are some secret managers which are planned to be implemented in future.")]),e._v(" "),a("ul",[a("li",[a("strong",[a("a",{attrs:{href:"https://azure.microsoft.com/en-in/services/key-vault/",target:"_blank",rel:"noopener noreferrer"}},[e._v("Azure Key Vault"),a("OutboundLink")],1)])]),e._v(" "),a("li",[a("strong",[a("a",{attrs:{href:"https://cloud.google.com/secret-manager",target:"_blank",rel:"noopener noreferrer"}},[e._v("GCP Secret Manager"),a("OutboundLink")],1)])])]),e._v(" "),a("h2",{attrs:{id:"supported-features"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#supported-features"}},[e._v("#")]),e._v(" Supported Features")]),e._v(" "),a("ul",[a("li",[e._v("Authentication to Hashicorp vault using Kubernetes service-account")]),e._v(" "),a("li",[e._v("RBAC implementation of vault using different policies of vault and association of policy with service-account")]),e._v(" "),a("li",[e._v("Inject secret directly to pods/containers running inside Kubernetes")]),e._v(" "),a("li",[e._v("Inject secret directly to pods/containers from AWS Secret Manager")]),e._v(" "),a("li",[e._v("Authentication with AWS Secret Manager with access key and iam role")]),e._v(" "),a("li",[e._v("Support regex to inject all secrets from a certain path of Vault")]),e._v(" "),a("li",[e._v("Inject secrets directly to the process of container, i.e. after the injection you cannot read secrets from the environment variable")])]),e._v(" "),a("h2",{attrs:{id:"architecture"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#architecture"}},[e._v("#")]),e._v(" Architecture")]),e._v(" "),a("h3",{attrs:{id:"hashicorp-vault"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#hashicorp-vault"}},[e._v("#")]),e._v(" Hashicorp Vault")]),e._v(" "),a("div",{staticStyle:{"padding-top":"25px"},attrs:{align:"center"}},[a("img",{attrs:{src:r(357)}})]),e._v(" "),a("h3",{attrs:{id:"aws-secret-manager"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#aws-secret-manager"}},[e._v("#")]),e._v(" AWS Secret Manager")]),e._v(" "),a("div",{staticStyle:{"padding-top":"25px"},attrs:{align:"center"}},[a("img",{attrs:{src:r(358)}})])])}),[],!1,null,null,null);t.default=n.exports}}]);