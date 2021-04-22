.. _development:

Pre-requisites
==============

**Access to Kubernetes cluster**

First of all, you will need access to a Kubernetes cluster. The easiest way to start is minikube.

- `Virtualbox <https://www.virtualbox.org/wiki/Downloads>`__ - hypervisor to run a Kubernetes cluster
- `Minikube <https://kubernetes.io/docs/setup/minikube/>`__ - for Kubernetes cluster creation on local machine
- `Kubectl <https://kubernetes.io/docs/tasks/tools/install-kubectl/>`__ - to interact with Kubernetes cluster


**Tools to build K8s Vault Webhook**

Apart from kubernetes cluster, there are some tools which are needed to build and test the vault k8s webhook.

- `Git <https://git-scm.com/downloads>`__
- `Go <https://golang.org/dl/>`__
- `Docker <https://docs.docker.com/install/>`__
- `Make <https://www.gnu.org/software/make/manual/make.html>`__

Build Locally
=============

To achieve this, execute this command:-

.. code:: bash

    $ make build-code

Build Image
===========

K8s Vault Webhook gets packaged as a container image for running on the Kubernetes cluster. These instructions will guide you to build an image.

.. code:: bash

    $ make build-image

Testing
=======

If you want to play it on Kubernetes. You can use a minikube.

.. code:: bash

    # Start minikube
    $ minikube start --vm-driver virtualbox

    # Deploy the image on minikube
    $ helm upgrade k8s-vault-webhook \
    ot-helm/k8s-vault-webhook --namespace vault --install

Run Tests
=========

.. code:: bash

    $ make test
