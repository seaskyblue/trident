.. _ontap-san-configuration-parameters:

###############################
ONTAP SAN Backend Configuration
###############################

========================= ==================================================================================================================================== ================================================
Parameter                 Description                                                                                                                          Default
========================= ==================================================================================================================================== ================================================
version                   Always 1
storageDriverName         "ontap-nas", "ontap-nas-economy", "ontap-nas-flexgroup", "ontap-san", "ontap-san-economy"
backendName               Custom name for the storage backend                                                                                                  Driver name + "_" + dataLIF
managementLIF             IP address of a cluster or SVM management LIF                                                                                        "10.0.0.1", "[2001:1234:abcd::fefe]"
dataLIF                   IP address of protocol LIF. **Use square brackets for IPv6**. Once set this **cannot be updated**                                    Derived by the SVM unless specified
useCHAP                   Use CHAP to authenticate iSCSI for ONTAP SAN drivers [Boolean]                                                                       false
chapInitiatorSecret       CHAP initiator secret. Required if ``useCHAP=true``                                                                                  ""
labels                    Set of arbitrary JSON-formatted labels to apply on volumes.                                                                          ""
chapTargetInitiatorSecret CHAP target initiator secret. Required if ``useCHAP=true``                                                                           ""
chapUsername              Inbound username. Required if ``useCHAP=true``                                                                                       ""
chapTargetUsername        Target username. Required if ``useCHAP=true``                                                                                        ""
clientCertificate         Base64-encoded value of client certificate. Used for certificate-based auth.                                                         ""
clientPrivateKey          Base64-encoded value of client private key. Used for certificate-based auth.                                                         ""
trustedCACertificate      Base64-encoded value of trusted CA certificate. Optional. Used for certificate-based auth.                                           ""
username                  Username to connect to the cluster/SVM. Used for credential-based auth.
password                  Password to connect to the cluster/SVM. Used for credential-based auth.
svm                       Storage virtual machine to use                                                                                                       Derived if an SVM managementLIF is specified
igroupName                Name of the igroup for SAN volumes to use                                                                                            "trident-<backend-UUID>"
username                  Username to connect to the cluster/SVM
password                  Password to connect to the cluster/SVM
storagePrefix             Prefix used when provisioning new volumes in the SVM. Once set this **cannot be updated**                                            "trident"
limitAggregateUsage       Fail provisioning if usage is above this percentage. **Does not apply to Amazon FSx for ONTAP**                                      "" (not enforced by default)
limitVolumeSize           Fail provisioning if requested volume size is above this value for the economy driver.                                               "" (not enforced by default)
lunsPerFlexvol            Maximum LUNs per Flexvol, must be in range [50, 200]                                                                                 "100"
debugTraceFlags           Debug flags to use when troubleshooting. E.g.: {"api":false, "method":true}                                                          null
========================= ==================================================================================================================================== ================================================

To communicate with the ONTAP cluster, Trident must be provided with authentication
parameters. This could be the username/password to a security login (OR) an
installed certificate. This is fully documented in the
:ref:`Authentication Guide <ontap-san-authentication>`.

.. warning::

  If you are using an Amazon FSx for ONTAP backend, do not specify the ``limitAggregateUsage`` parameter. The ``fsxadmin`` and ``vsadmin`` roles provided by Amazon FSx for ONTAP do not contain the required access permissions to retrieve aggregate usage and limit it through Trident.

.. warning::

  Do not use ``debugTraceFlags`` unless you are troubleshooting and require a
  detailed log dump.

For the ``ontap-san*`` drivers, the default is to use all data LIF IPs from
the SVM and to use iSCSI multipath. Specifying an IP address for the ``dataLIF``
for the ``ontap-san*`` drivers forces them to disable multipath and use only the
specified address.

.. note::

   When creating a backend, remember that the ``dataLIF`` and ``storagePrefix``
   cannot be modified after creation. To update these parameters you will need
   to create a new backend.

The ``igroupName`` can be set to an igroup that is already created on the ONTAP cluster.
If unspecified, Trident automatically creates an igroup named ``trident-<backend-UUID>``.
If providing a pre-defined ``igroupName``, NetApp recommends using an igroup per
Kubernetes cluster, if the SVM is to be shared between environments. This is
necessary for Trident to maintain IQN additions/deletions automatically.

Backends can also have igroups updated after creation:

* ``igroupName`` can be updated to point to a new igroup that is created and managed
  on the SVM outside of Trident.
* ``igroupName`` can be omitted. In this case, Trident will create and manage a
  ``trident-<backend-UUID>`` igroup automatically.

In both cases, volume attachments will continue to be accessible.
Future volume attachments will use the updated igroup. This update does not disrupt
access to volumes present on the backend.

A fully-qualified domain name (FQDN) can be specified for the ``managementLIF``
option.

The ``managementLIF`` for all ONTAP drivers can
also be set to IPv6 addresses. Make sure to install Trident with the
``--use-ipv6`` flag. Care must be taken to define the ``managementLIF``
IPv6 address **within square brackets**.

.. warning::

   When using IPv6 addresses, make sure the ``managementLIF`` and ``dataLIF``
   [if included in your backend defition] are defined
   within square brackets, such as ``[28e8:d9fb:a825:b7bf:69a8:d02f:9e7b:3555]``.
   If the ``dataLIF`` is not provided, Trident will fetch the IPv6 data LIFs
   from the SVM.

To enable the ``ontap-san*`` drivers to use CHAP, set the ``useCHAP`` parameter to
``true`` in your backend definition. Trident will then configure and use
bidirectional CHAP as the default authentication for the SVM given in the backend.
The :ref:`CHAP with ONTAP SAN drivers<Using CHAP with ONTAP SAN drivers>`
section explains how this works.

For the ``ontap-san-economy`` driver, the ``limitVolumeSize``
option will also restrict the maximum size of
the volumes it manages for qtrees and LUNs.

.. note::

  Trident sets provisioning labels in the "Comments" field of all volumes
  created using the ``ontap-san`` driver. For each volume created, the "Comments"
  field on the FlexVol will be populated with all labels present on the storage
  pool it is placed in. Storage admins can define labels per storage pool and
  group all volumes created in a storage pool. This provides a convenient way of
  differentiating volumes based on a set of customizable labels that are
  provided in the backend configuration.

Configuration options for provisioning volumes
----------------------------------------------

You can control how each volume is provisioned by default using these options
in a special section of the configuration. For an example, see the
configuration examples below.

========================= =============================================================== ================================================
Parameter                 Description                                                     Default
========================= =============================================================== ================================================
spaceAllocation           Space-allocation for LUNs                                       "true"
spaceReserve              Space reservation mode; "none" (thin) or "volume" (thick)       "none"
snapshotPolicy            Snapshot policy to use                                          "none"
qosPolicy                 QoS policy group to assign for volumes created.
                          Choose one of ``qosPolicy`` or ``adaptiveQosPolicy`` per
                          storage pool/backend.                                           ""
adaptiveQosPolicy         Adaptive QoS policy group to assign for volumes created.
                          Choose one of ``qosPolicy`` or
                          ``adaptiveQosPolicy`` per storage pool/backend.                 ""
snapshotReserve           Percentage of volume reserved for snapshots                     "0" if snapshotPolicy is "none", else ""
splitOnClone              Split a clone from its parent upon creation                     "false"
encryption                Enable NetApp volume encryption                                 "false"
securityStyle             Security style for new volumes                                  "unix"
tieringPolicy             Tiering policy to use                                           "none"; "snapshot-only" for pre-ONTAP 9.5 SVM-DR configuration
========================= =============================================================== ================================================

.. note::

  Using QoS policy groups with Trident requires ONTAP 9.8 or later.
  It is recommended to use a **non-shared** QoS policy group and ensure the policy
  group is applied to each constituent **individually**. A shared QoS policy group
  will enforce the ceiling for the **total throughput** of all workloads.

Here's an example with defaults defined:

.. code-block:: bash

  {
   "version": 1,
   "storageDriverName": "ontap-san",
   "managementLIF": "10.0.0.1",
   "dataLIF": "10.0.0.2",
   "svm": "trident_svm",
   "username": "admin",
   "password": "password",
   "labels": {"k8scluster": "dev2", "backend": "dev2-sanbackend"},
   "storagePrefix": "alternate-trident",
   "igroupName": "custom",
   "debugTraceFlags": {"api":false, "method":true},
   "defaults": {
       "spaceReserve": "volume",
       "qosPolicy": "standard",
       "spaceAllocation": "false",
       "snapshotPolicy": "default",
       "snapshotReserve": "10"
   }
  }

.. note::

   For all volumes created using the ``ontap-san`` driver, Trident adds an extra 10 percent capacity to the FlexVol to accommodate the LUN metadata. The LUN will be provisioned with the exact size that the user requests in the PVC. Trident adds 10 percent to the FlexVol (shows as ``Available`` size in ONTAP). Users will now get the amount of usable capacity they requested. This change also prevents LUNs from becoming read-only unless the available space is fully utilized. This **does not** apply to ``ontap-san-economy``.

For backends that define ``snapshotReserve``, Trident calculates the size of volumes as follows:

Total volume size = [(PVC requested size) / (1 - (``snapshotReserve`` percentage) / 100)] * 1.1

The 1.1 is the extra 10 percent Trident adds to the FlexVol to accommodate the LUN metadata. For ``snapshotReserve`` = 5%, and PVC request = 5GiB, the total volume size is 5.79GiB and the available size is 5.5GiB. The ``volume show`` command should show results similar to this example:

.. _figVolshow:

.. figure:: images/vol-show-san.png
    :align: center
    :figclass: alight-center

Currently, resizing is the only way to use the new calculation for an existing volume.
